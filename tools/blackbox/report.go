package main

import (
	"fmt"
	"io"
	"math"
	"sort"
	"strings"
	"text/tabwriter"
	"time"
)

type procKey struct {
	pid  int
	comm string
}

type procAgg struct {
	PID    int
	Comm   string
	CPU    float64 // cpu-seconds
	BlkIO  float64 // seconds waiting on block I/O
	DSec   float64 // thread-seconds observed in D state
	Rd     int64
	Wr     int64
	CWr    int64
	MajFlt int64
	MaxRSS int64 // kB
}

type diskAgg struct {
	rkb, wkb, utilSum, dt, maxUtil, maxAwait float64
}

type incident struct {
	start, end int64 // end is 0 while still open
	reason     string
}

// report aggregates records over a time window.
type report struct {
	first, last int64
	samples     int
	procs       map[procKey]*procAgg
	cmds        map[int]string

	dTotal, dFault                 float64
	dThread, wchan, sysfile, stack map[string]float64

	peakLoad              float64
	peakRunq, peakBlk     int
	peakD                 int
	iowSum, iowMax, sumDt float64
	psiSum, psiMax        float64
	hasPSI                bool
	minAvail, memTotal    int64
	maxSwap, maxDirty     int64
	majflt, swapKB        float64
	oom, dropped          uint64
	anyBlkIO              bool
	disks                 map[string]*diskAgg
	incidents             []incident
}

func newReport() *report {
	return &report{
		procs:    map[procKey]*procAgg{},
		cmds:     map[int]string{},
		dThread:  map[string]float64{},
		wchan:    map[string]float64{},
		sysfile:  map[string]float64{},
		stack:    map[string]float64{},
		disks:    map[string]*diskAgg{},
		minAvail: -1,
	}
}

func (r *report) proc(pid int, comm string) *procAgg {
	k := procKey{pid, comm}
	p := r.procs[k]
	if p == nil {
		p = &procAgg{PID: pid, Comm: comm}
		r.procs[k] = p
	}
	return p
}

func (r *report) add(rec *Record) {
	if r.first == 0 {
		r.first = rec.T
	}
	r.last = rec.T
	r.dropped += rec.Dropped
	switch rec.Kind {
	case "sample":
		r.addSample(rec)
	case "blk":
		d := rec.Dt
		if d <= 0 {
			d = 1
		}
		if len(rec.Blocked) > r.peakD {
			r.peakD = len(rec.Blocked)
		}
		for _, b := range rec.Blocked {
			r.dTotal += d
			r.proc(b.PID, b.Comm).DSec += d
			name := b.Comm
			if b.TComm != "" {
				name += "/" + b.TComm
			}
			r.dThread[name] += d
			w := b.Wchan
			if w == "" {
				w = "?"
			}
			r.wchan[w] += d
			sf := b.Syscall
			if sf == "" {
				sf = "(kernel thread)"
			}
			if b.File != "" {
				sf += "  " + b.File
			}
			r.sysfile[b.Comm+": "+sf] += d
			if b.Syscall == "pagefault" {
				r.dFault += d
			}
			if len(b.Stack) > 0 {
				r.stack[strings.Join(b.Stack, "\n")] += d
			}
		}
	case "inc":
		switch rec.Event {
		case "start":
			r.incidents = append(r.incidents, incident{start: rec.T, reason: rec.Reason})
		case "end":
			if n := len(r.incidents); n > 0 && r.incidents[n-1].end == 0 {
				r.incidents[n-1].end = rec.T
			} else {
				// Start fell before the window.
				r.incidents = append(r.incidents, incident{start: rec.T - int64(rec.Dt*1000), end: rec.T, reason: "(started before window)"})
			}
		}
	case "boot":
		// A restart closes any incident left open by the previous run.
		if n := len(r.incidents); n > 0 && r.incidents[n-1].end == 0 {
			r.incidents[n-1].end = rec.T
			r.incidents[n-1].reason += " (recorder restarted)"
		}
	}
}

func (r *report) addSample(rec *Record) {
	r.samples++
	dt := rec.Dt
	s := rec.Sys
	if s != nil {
		r.sumDt += dt
		r.peakLoad = math.Max(r.peakLoad, s.Load1)
		r.peakRunq = max(r.peakRunq, s.Running)
		r.peakBlk = max(r.peakBlk, s.Blocked)
		r.iowSum += s.IOWait * dt
		r.iowMax = math.Max(r.iowMax, s.IOWait)
		if s.PSI != nil {
			r.hasPSI = true
			r.psiSum += s.PSI.IOSome * dt
			r.psiMax = math.Max(r.psiMax, s.PSI.IOSome)
		}
		if r.minAvail < 0 || s.MemAvail < r.minAvail {
			r.minAvail = s.MemAvail
		}
		r.memTotal = s.MemTotal
		r.maxSwap = max(r.maxSwap, s.SwapUsed)
		r.maxDirty = max(r.maxDirty, s.Dirty+s.Writeback)
		r.majflt += s.MajFlt * dt
		r.swapKB += (s.SwpIn + s.SwpOut) * dt * 4
		r.oom += s.OOMKill
		for _, d := range s.Disks {
			a := r.disks[d.Name]
			if a == nil {
				a = &diskAgg{}
				r.disks[d.Name] = a
			}
			a.rkb += d.RKB * dt
			a.wkb += d.WKB * dt
			a.utilSum += d.Util * dt
			a.maxUtil = math.Max(a.maxUtil, d.Util)
			a.maxAwait = math.Max(a.maxAwait, d.Await)
		}
	}
	for i := range rec.Procs {
		p := &rec.Procs[i]
		if p.Cmd != "" {
			r.cmds[p.PID] = p.Cmd
		}
		a := r.proc(p.PID, p.Comm)
		a.CPU += p.CPU / 100 * dt
		a.BlkIO += float64(p.BlkIO) / 1000
		a.Rd += p.RdB
		a.Wr += p.WrB
		a.CWr += p.CWrB
		a.MajFlt += p.MajFlt
		a.MaxRSS = max(a.MaxRSS, p.RSS)
		if p.BlkIO > 0 {
			r.anyBlkIO = true
		}
	}
}

// ---- formatting helpers ----

func fmtTime(ms int64) string { return time.UnixMilli(ms).Local().Format("01-02 15:04:05") }

// hb formats a byte count.
func hb(n float64) string {
	const u = "BKMGT"
	i := 0
	for math.Abs(n) >= 1024 && i < len(u)-1 {
		n /= 1024
		i++
	}
	if i == 0 {
		return fmt.Sprintf("%.0f", n)
	}
	return fmt.Sprintf("%.1f%c", n, u[i])
}

func hkb(kb int64) string { return hb(float64(kb) * 1024) }

func secs(s float64) string {
	if s == 0 {
		return "-"
	}
	return fmt.Sprintf("%.1f", s)
}

type kv struct {
	k string
	v float64
}

func topOf(m map[string]float64, n int) []kv {
	out := make([]kv, 0, len(m))
	for k, v := range m {
		out = append(out, kv{k, v})
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].v != out[j].v {
			return out[i].v > out[j].v
		}
		return out[i].k < out[j].k
	})
	if len(out) > n {
		out = out[:n]
	}
	return out
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n-1] + "…"
}

// ---- report sections ----

func (r *report) printSummary(w io.Writer) {
	fmt.Fprintf(w, "Window %s .. %s, %d samples\n", fmtTime(r.first), fmtTime(r.last), r.samples)
	if r.sumDt > 0 {
		fmt.Fprintf(w, "  load1 peak %.2f   run queue peak %d   procs_blocked peak %d   D threads peak %d\n",
			r.peakLoad, r.peakRunq, r.peakBlk, r.peakD)
		fmt.Fprintf(w, "  iowait avg %.1f%% max %.1f%%", r.iowSum/r.sumDt, r.iowMax)
		if r.hasPSI {
			fmt.Fprintf(w, "   psi io(some) avg %.1f%% max %.1f%%", r.psiSum/r.sumDt, r.psiMax)
		}
		fmt.Fprintln(w)
		fmt.Fprintf(w, "  MemAvailable min %s of %s   dirty+writeback max %s   swap used max %s\n",
			hkb(r.minAvail), hkb(r.memTotal), hkb(r.maxDirty), hkb(r.maxSwap))
		fmt.Fprintf(w, "  major faults %.0f   swap I/O %s", r.majflt, hkb(int64(r.swapKB)))
		if r.oom > 0 {
			fmt.Fprintf(w, "   OOM kills %d", r.oom)
		}
		fmt.Fprintln(w)
	}
	if r.dropped > 0 {
		fmt.Fprintf(w, "  %d records were dropped because the disk writer fell behind\n", r.dropped)
	}
	if len(r.disks) > 0 {
		tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', tabwriter.AlignRight)
		fmt.Fprintln(tw, "  DISK\tREAD\tWRITTEN\tUTIL avg\tUTIL max\tAWAIT max ms\t")
		names := make([]string, 0, len(r.disks))
		for n := range r.disks {
			names = append(names, n)
		}
		sort.Slice(names, func(i, j int) bool { return r.disks[names[i]].utilSum > r.disks[names[j]].utilSum })
		for i, n := range names {
			if i == 8 {
				break
			}
			d := r.disks[n]
			fmt.Fprintf(tw, "  %s\t%s\t%s\t%.1f%%\t%.1f%%\t%.1f\t\n", n, hkb(int64(d.rkb)), hkb(int64(d.wkb)),
				d.utilSum/math.Max(r.sumDt, 1e-9), d.maxUtil, d.maxAwait)
		}
		tw.Flush()
	}
}

var sortKeys = map[string]func(p *procAgg) float64{
	"blkio":  func(p *procAgg) float64 { return p.BlkIO },
	"dstate": func(p *procAgg) float64 { return p.DSec },
	"cpu":    func(p *procAgg) float64 { return p.CPU },
	"read":   func(p *procAgg) float64 { return float64(p.Rd) },
	"write":  func(p *procAgg) float64 { return float64(p.Wr) },
	"io":     func(p *procAgg) float64 { return float64(p.Rd + p.Wr) },
	"majflt": func(p *procAgg) float64 { return float64(p.MajFlt) },
	"rss":    func(p *procAgg) float64 { return float64(p.MaxRSS) },
	// Default ranking: anything that contributed to I/O stalls.
	"stall": func(p *procAgg) float64 { return p.BlkIO + p.DSec + float64(p.Rd+p.Wr)/(16<<20) },
}

func (r *report) printProcs(w io.Writer, by string, n int) {
	key := sortKeys[by]
	list := make([]*procAgg, 0, len(r.procs))
	for _, p := range r.procs {
		if key(p) > 0 {
			list = append(list, p)
		}
	}
	sort.Slice(list, func(i, j int) bool { return key(list[i]) > key(list[j]) })
	if len(list) > n {
		list = list[:n]
	}
	fmt.Fprintf(w, "\nProcesses by %s\n", by)
	if len(list) == 0 {
		fmt.Fprintln(w, "  (none)")
		return
	}
	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)
	fmt.Fprintln(tw, "  PID\tCOMM\tBLKIO s\tD thr·s\tREAD\tWRITE\tCANCEL\tMAJFLT\tCPU s\tRSS max\tCOMMAND")
	for _, p := range list {
		fmt.Fprintf(tw, "  %d\t%s\t%s\t%s\t%s\t%s\t%s\t%d\t%s\t%s\t%s\n", p.PID, p.Comm, secs(p.BlkIO), secs(p.DSec),
			hb(float64(p.Rd)), hb(float64(p.Wr)), hb(float64(p.CWr)), p.MajFlt, secs(p.CPU), hkb(p.MaxRSS),
			truncate(r.cmds[p.PID], 60))
	}
	tw.Flush()
}

func (r *report) printBlocked(w io.Writer, n, stacks int) {
	fmt.Fprintf(w, "\nD-state (uninterruptible) threads: %.1f thread-seconds observed\n", r.dTotal)
	if r.dTotal == 0 {
		fmt.Fprintln(w, "  (none)")
		return
	}
	section := func(title string, m map[string]float64) {
		fmt.Fprintf(w, "  %s\n", title)
		for _, e := range topOf(m, n) {
			fmt.Fprintf(w, "    %7.1f s  %4.0f%%  %s\n", e.v, e.v*100/r.dTotal, e.k)
		}
	}
	section("by thread (process/thread):", r.dThread)
	section("by kernel wait point (wchan):", r.wchan)
	section("by syscall and file:", r.sysfile)
	if stacks > 0 && len(r.stack) > 0 {
		fmt.Fprintln(w, "  most common kernel stacks:")
		for _, e := range topOf(r.stack, stacks) {
			fmt.Fprintf(w, "    %.1f s:\n", e.v)
			for _, f := range strings.Split(e.k, "\n") {
				fmt.Fprintf(w, "        %s\n", f)
			}
		}
	}
}

// printHints prints conclusions that follow directly from the data.
func (r *report) printHints(w io.Writer) {
	var hints []string
	if r.dTotal > 0 && r.dFault/r.dTotal >= 0.3 {
		hints = append(hints, fmt.Sprintf("%.0f%% of blocked time was spent in page faults, not in file syscalls. "+
			"Memory is short enough that program code and mapped files are being evicted and re-read from storage "+
			"(page-cache thrashing). Compare MemAvailable and major faults, and look at the largest RSS.", r.dFault*100/r.dTotal))
	}
	if r.swapKB > 10*1024 {
		hints = append(hints, fmt.Sprintf("%s of swap I/O in this window; swapping competes with everything else for the disk.", hkb(int64(r.swapKB))))
	}
	var jbd, throttle float64
	for k, v := range r.stack {
		if strings.Contains(k, "jbd2") {
			jbd += v
		}
		if strings.Contains(k, "balance_dirty_pages") {
			throttle += v
		}
	}
	for k, v := range r.wchan {
		if strings.Contains(k, "jbd2") {
			jbd += v
		}
		if strings.Contains(k, "balance_dirty_pages") {
			throttle += v
		}
	}
	if r.dTotal > 0 && jbd/r.dTotal >= 0.2 {
		hints = append(hints, "Many threads waited on the ext4 journal (jbd2). This usually comes from processes calling "+
			"fsync/fdatasync often; the fsync callers show up under 'by syscall and file', and the journal itself "+
			"is written by the jbd2/<device> kernel thread.")
	}
	if r.dTotal > 0 && throttle/r.dTotal >= 0.2 {
		hints = append(hints, "Writers were throttled in balance_dirty_pages: dirty data is produced faster than the "+
			"disk can write it back. The top writers by WRITE are the source.")
	}
	var totalWr, topWr int64
	var topName string
	for _, p := range r.procs {
		totalWr += p.Wr
		if p.Wr > topWr {
			topWr, topName = p.Wr, fmt.Sprintf("%s (pid %d)", p.Comm, p.PID)
		}
	}
	if totalWr > 16<<20 && float64(topWr) >= 0.5*float64(totalWr) {
		hints = append(hints, fmt.Sprintf("%s dirtied %.0f%% (%s) of all bytes written by processes in this window.",
			topName, float64(topWr)*100/float64(totalWr), hb(float64(topWr))))
	}
	if r.samples > 0 && !r.anyBlkIO && r.dTotal > 0 {
		hints = append(hints, "BLKIO is zero for every process although threads were blocked: per-task delay accounting "+
			"is off. Run 'sysctl -w kernel.task_delayacct=1' or start the recorder with -delayacct.")
	}
	if len(hints) == 0 {
		return
	}
	fmt.Fprintln(w, "\nNotes")
	for _, h := range hints {
		fmt.Fprintf(w, "  - %s\n", h)
	}
}

// ---- live/timeline views ----

const sysHeader = "TIME            LOAD1  RUN  BLK   D   USR  SYS  IOW  IDLE  PSIio  AVAIL  DIRTY   SWAP  MAJF/s  DISKS"

func sysLine(t int64, s *Sys, dThreads int) string {
	psiIO := "-"
	if s.PSI != nil {
		psiIO = fmt.Sprintf("%.1f", s.PSI.IOSome)
	}
	disks := make([]Disk, len(s.Disks))
	copy(disks, s.Disks)
	sort.Slice(disks, func(i, j int) bool { return disks[i].Util > disks[j].Util })
	var ds []string
	for i, d := range disks {
		if i == 2 {
			break
		}
		ds = append(ds, fmt.Sprintf("%s %.0f%% r%s w%s", d.Name, d.Util, hb(d.RKB*1024), hb(d.WKB*1024)))
	}
	return fmt.Sprintf("%s  %5.2f  %3d  %3d %3d  %4.0f %4.0f %4.0f  %4.0f  %5s  %5s  %5s  %5s  %6.0f  %s",
		fmtTime(t), s.Load1, s.Running, s.Blocked, dThreads, s.User+s.Nice, s.System+s.IRQ+s.SoftIRQ,
		s.IOWait, s.Idle, psiIO, hkb(s.MemAvail), hkb(s.Dirty+s.Writeback), hkb(s.SwapUsed), s.MajFlt,
		strings.Join(ds, ", "))
}

func printProcRows(w io.Writer, procs []Proc, dt float64, n int, indent string) {
	ps := make([]Proc, len(procs))
	copy(ps, procs)
	sort.Slice(ps, func(i, j int) bool { return procScore(&ps[i]) > procScore(&ps[j]) })
	if len(ps) > n {
		ps = ps[:n]
	}
	if dt <= 0 {
		dt = 1
	}
	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)
	fmt.Fprintf(tw, "%sPID\tCOMM\tST\tCPU%%\tBLKIO ms\tREAD/s\tWRITE/s\tMAJF\tRSS\tCOMMAND\n", indent)
	for _, p := range ps {
		fmt.Fprintf(tw, "%s%d\t%s\t%s\t%.1f\t%d\t%s\t%s\t%d\t%s\t%s\n", indent, p.PID, p.Comm, p.State, p.CPU, p.BlkIO,
			hb(float64(p.RdB)/dt), hb(float64(p.WrB)/dt), p.MajFlt, hkb(p.RSS), truncate(p.Cmd, 50))
	}
	tw.Flush()
}

func printBlockedRows(w io.Writer, blk []Blocked, withStacks bool) {
	for _, b := range blk {
		name := b.Comm
		if b.TComm != "" {
			name += "/" + b.TComm
		}
		fmt.Fprintf(w, "  %d/%d %s  wchan=%s  syscall=%s", b.PID, b.TID, name, orDash(b.Wchan), orDash(b.Syscall))
		if b.File != "" {
			fmt.Fprintf(w, "  file=%s", b.File)
		}
		fmt.Fprintln(w)
		if withStacks {
			for _, f := range b.Stack {
				fmt.Fprintf(w, "      %s\n", f)
			}
		}
	}
}

func orDash(s string) string {
	if s == "" {
		return "-"
	}
	return s
}
