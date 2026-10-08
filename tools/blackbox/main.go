// Command blackbox is a low-overhead flight recorder for small Linux systems
// such as UniFi gateways. It samples /proc continuously, keeps a bounded,
// compressed history on disk, and when load or I/O pressure crosses a
// threshold it records which threads are stuck in uninterruptible sleep,
// where in the kernel, and on which file.
package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"runtime/debug"
	"strings"
	"syscall"
	"text/tabwriter"
	"time"
)

const usage = `blackbox - low-overhead flight recorder for small Linux systems

Usage:
  blackbox record    [flags]  run the recorder in the foreground
  blackbox now       [flags]  live view: top processes and blocked threads right now
  blackbox incidents [flags]  list load/I-O incidents with a summary of each
  blackbox top       [flags]  per-process totals over a time window
  blackbox blocked   [flags]  D-state threads over a time window: who, where, which file
  blackbox show      [flags]  system timeline, one line per sample
  blackbox export    [flags]  dump raw records as JSON lines
  blackbox check              report which kernel features are available

Time flags (-since, -until) take a duration ago ("90m"), a clock time today
("14:05"), or a date and time ("2026-10-08 14:05").

Run 'blackbox <command> -h' for the flags of each command.
`

const defaultDir = "/data/blackbox/data"

func main() {
	log.SetFlags(log.LstdFlags)
	if len(os.Args) < 2 {
		fmt.Fprint(os.Stderr, usage)
		os.Exit(2)
	}
	cmds := map[string]func([]string) error{
		"record": cmdRecord, "now": cmdNow, "incidents": cmdIncidents, "top": cmdTop,
		"blocked": cmdBlocked, "show": cmdShow, "export": cmdExport, "check": cmdCheck,
	}
	name := os.Args[1]
	fn, ok := cmds[name]
	if !ok {
		if name == "-h" || name == "--help" || name == "help" {
			fmt.Print(usage)
			return
		}
		fmt.Fprintf(os.Stderr, "unknown command %q\n\n%s", name, usage)
		os.Exit(2)
	}
	if err := fn(os.Args[2:]); err != nil {
		fmt.Fprintln(os.Stderr, "blackbox:", err)
		os.Exit(1)
	}
}

// ---- record ----

func cmdRecord(args []string) error {
	fs := flag.NewFlagSet("record", flag.ExitOnError)
	dir := fs.String("dir", defaultDir, "data directory")
	interval := fs.Duration("interval", 5*time.Second, "system and process sample interval")
	incInterval := fs.Duration("incident-interval", 2*time.Second, "sample interval during an incident")
	probe := fs.Duration("probe", time.Second, "how often to check load and look for blocked threads")
	incProbe := fs.Duration("incident-probe", 500*time.Millisecond, "blocked-thread scan interval during an incident")
	fullEvery := fs.Duration("full-every", time.Minute, "interval between snapshots that list every process")
	topN := fs.Int("top", 40, "max processes per regular sample (only active ones are recorded)")
	load := fs.Float64("load", 0, "load1 at or above this starts an incident (0 = number of CPUs)")
	psiIO := fs.Float64("psi-io", 20, "PSI io 'some' avg10 % at or above this starts an incident (0 = off)")
	blocked := fs.Int("blocked", 2, "procs_blocked at or above this for 3 probes in a row starts an incident")
	hold := fs.Duration("hold", 30*time.Second, "an incident ends after this long with no trigger")
	maxDetail := fs.Duration("max-detail", 10*time.Minute, "stop fast scans and stack capture after an incident has lasted this long")
	maxBlocked := fs.Int("max-blocked", 64, "max D-state threads recorded per scan")
	frames := fs.Int("stack-frames", 16, "max kernel stack frames kept per blocked thread")
	flushEvery := fs.Duration("flush", 30*time.Second, "flush buffered data to disk this often (also on every incident sample)")
	segAge := fs.Duration("segment-age", time.Hour, "start a new data file after this long")
	segMB := fs.Int64("segment-mb", 4, "start a new data file after this many compressed MB")
	maxMB := fs.Int64("max-mb", 64, "total size of the data directory; oldest files are deleted")
	memLimit := fs.Int64("mem-limit-mb", 12, "soft limit for the Go heap")
	delayacct := fs.Bool("delayacct", false, "turn on kernel.task_delayacct at start if it is off (needed for per-process BLKIO)")
	mlock := fs.Bool("mlock", false, "lock the program's code in RAM so memory pressure cannot evict it (~2.5 MB)")
	fs.Parse(args)

	runtime.GOMAXPROCS(1)
	debug.SetGCPercent(50)
	debug.SetMemoryLimit(*memLimit << 20)

	if *delayacct {
		if b, err := os.ReadFile("/proc/sys/kernel/task_delayacct"); err == nil && strings.TrimSpace(string(b)) == "0" {
			if err := os.WriteFile("/proc/sys/kernel/task_delayacct", []byte("1\n"), 0); err != nil {
				log.Printf("enable task_delayacct: %v", err)
			} else {
				log.Printf("enabled kernel.task_delayacct")
			}
		}
	}
	if *mlock {
		if n, err := lockExecutable(); err != nil {
			log.Printf("mlock: %v", err)
		} else {
			log.Printf("locked %d kB of program text and data in memory", n>>10)
		}
	}

	st, err := OpenStore(StoreConfig{
		Dir: *dir, FlushEvery: *flushEvery, SegMaxAge: *segAge,
		SegMaxBytes: *segMB << 20, MaxTotal: *maxMB << 20,
	})
	if err != nil {
		return err
	}
	col := NewCollector(Config{
		Interval: *interval, FastInterval: *probe, IncidentInterval: *incProbe,
		FullEvery: *fullEvery, TopN: *topN, LoadThreshold: *load, PSIThreshold: *psiIO,
		BlockedThreshold: *blocked, IncidentHold: *hold, MaxDetail: *maxDetail,
		MaxBlocked: *maxBlocked, StackFrames: *frames,
	}, st.Put)

	st.Put(&Record{T: time.Now().UnixMilli(), Kind: "boot", Boot: bootInfo()})
	log.Printf("recording to %s (load threshold %.1f)", *dir, col.cfg.LoadThreshold)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	now := time.Now()
	col.Sample(now)
	nextSample := now.Add(*interval)
	timer := time.NewTimer(0)
	<-timer.C
	for {
		now = time.Now()
		inIncident := col.Probe(now)
		if !now.Before(nextSample) {
			if rec := col.Sample(now); rec != nil {
				rec.Dropped = st.TakeDropped()
				st.Put(rec)
			}
			if inIncident {
				nextSample = now.Add(*incInterval)
				st.FlushSoon()
			} else {
				nextSample = now.Add(*interval)
			}
		}
		wait := *probe
		if col.Detailed(now) {
			wait = *incProbe
		}
		timer.Reset(wait)
		select {
		case s := <-sig:
			log.Printf("%v: flushing and exiting", s)
			col.Stop(time.Now())
			st.Close()
			return nil
		case <-timer.C:
		}
	}
}

// lockExecutable mlocks the mappings of the program binary. Under memory
// pressure the kernel evicts code pages, and a recorder that has to page its
// own code back in from a saturated disk stalls exactly when it is needed.
// mlockall would also lock the Go runtime's address-space reservations,
// which costs tens of MB.
func lockExecutable() (int, error) {
	exe, err := os.Readlink("/proc/self/exe")
	if err != nil {
		return 0, err
	}
	maps, err := os.ReadFile("/proc/self/maps")
	if err != nil {
		return 0, err
	}
	total := 0
	for _, line := range strings.Split(string(maps), "\n") {
		f := strings.Fields(line)
		if len(f) < 6 || f[5] != exe {
			continue
		}
		var lo, hi uintptr
		if _, err := fmt.Sscanf(f[0], "%x-%x", &lo, &hi); err != nil {
			continue
		}
		if _, _, e := syscall.Syscall(syscall.SYS_MLOCK, lo, hi-lo, 0); e != 0 {
			return total, e
		}
		total += int(hi - lo)
	}
	return total, nil
}

func bootInfo() *Boot {
	host, _ := os.Hostname()
	var u syscall.Utsname
	kernel := ""
	if syscall.Uname(&u) == nil {
		kernel = utsString(u.Release[:])
	}
	return &Boot{Host: host, Kernel: kernel, Arch: runtime.GOARCH, NCPU: runtime.NumCPU(), Features: features()}
}

func utsString[T int8 | uint8](b []T) string {
	out := make([]byte, 0, len(b))
	for _, c := range b {
		if c == 0 {
			break
		}
		out = append(out, byte(c))
	}
	return string(out)
}

// features probes the kernel interfaces blackbox depends on.
func features() map[string]string {
	f := map[string]string{}
	ok := func(path string) string {
		if _, err := os.ReadFile(path); err != nil {
			var pe *os.PathError
			if errors.As(err, &pe) {
				err = pe.Err
			}
			return "no (" + err.Error() + ")"
		}
		return "yes"
	}
	f["io_accounting"] = ok("/proc/self/io")
	f["psi"] = ok("/proc/pressure/io")
	f["stack"] = ok("/proc/self/stack")
	f["syscall"] = ok("/proc/self/syscall")
	f["wchan"] = ok("/proc/self/wchan")
	if b, err := os.ReadFile("/proc/sys/kernel/task_delayacct"); err == nil {
		f["task_delayacct"] = strings.TrimSpace(string(b))
	} else {
		f["task_delayacct"] = "no sysctl (kernel < 5.14: on if CONFIG_TASK_DELAY_ACCT=y)"
	}
	return f
}

// ---- check ----

func cmdCheck(args []string) error {
	f := features()
	var u syscall.Utsname
	syscall.Uname(&u)
	fmt.Printf("kernel %s, %s, %d CPUs, page size %d\n\n", utsString(u.Release[:]), runtime.GOARCH, runtime.NumCPU(), os.Getpagesize())
	explain := []struct{ key, what string }{
		{"io_accounting", "per-process bytes read/written (/proc/<pid>/io, CONFIG_TASK_IO_ACCOUNTING)"},
		{"task_delayacct", "per-process time spent waiting on block I/O (BLKIO column)"},
		{"psi", "pressure stall information (CONFIG_PSI; may need psi=1 on the kernel command line)"},
		{"stack", "kernel stacks of blocked threads (CONFIG_STACKTRACE, root)"},
		{"syscall", "syscall a blocked thread is in, used to find the file (CONFIG_HAVE_ARCH_TRACEHOOK)"},
		{"wchan", "kernel function a blocked thread is waiting in"},
	}
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	for _, e := range explain {
		fmt.Fprintf(tw, "%s\t%s\t%s\n", e.key, f[e.key], e.what)
	}
	tw.Flush()

	// Delay accounting can be compiled in but off; see whether any process
	// has accumulated block I/O delay.
	r := newReader()
	var ticks uint64
	for _, pid := range listNumeric("/proc") {
		if b, err := r.read(fmt.Sprintf("/proc/%d/stat", pid)); err == nil {
			if ps, ok := parsePidStat(b); ok {
				ticks += ps.blkio
			}
		}
	}
	fmt.Printf("\nsum of delayacct_blkio_ticks over all processes: %d", ticks)
	if ticks == 0 {
		fmt.Print("  (zero: delay accounting is off or not compiled in; try 'sysctl -w kernel.task_delayacct=1')")
	}
	fmt.Println()
	if os.Geteuid() != 0 {
		fmt.Println("\nnot running as root: other users' /proc/<pid>/io, stack and fd links will be unreadable")
	}
	return nil
}

// ---- now ----

func cmdNow(args []string) error {
	fs := flag.NewFlagSet("now", flag.ExitOnError)
	n := fs.Int("n", 15, "processes to show")
	iv := fs.Duration("i", time.Second, "measurement interval")
	stacks := fs.Bool("stacks", true, "show kernel stacks of blocked threads")
	fs.Parse(args)

	col := NewCollector(Config{FullEvery: 1<<62 - 1, TopN: 1 << 20, MaxBlocked: 256, StackFrames: 16}, func(*Record) {})
	col.Sample(time.Now())
	time.Sleep(*iv)
	rec := col.Sample(time.Now())
	blk := col.ScanBlocked(*stacks)

	fmt.Println(sysHeader)
	fmt.Println(sysLine(rec.T, rec.Sys, len(blk)))
	if p := rec.Sys.PSI; p != nil {
		fmt.Printf("PSI: cpu %.1f%%  memory some %.1f%% full %.1f%%  io some %.1f%% full %.1f%%\n",
			p.CPU, p.MemSome, p.MemFull, p.IOSome, p.IOFull)
	}
	fmt.Println()
	printProcRows(os.Stdout, rec.Procs, rec.Dt, *n, "")
	fmt.Printf("\n%d threads in D state\n", len(blk))
	printBlockedRows(os.Stdout, blk, *stacks)
	return nil
}

// ---- readers ----

type windowFlags struct {
	dir, since, until *string
}

func addWindowFlags(fs *flag.FlagSet, defSince string) windowFlags {
	return windowFlags{
		dir:   fs.String("dir", defaultDir, "data directory"),
		since: fs.String("since", defSince, "start of window (empty = beginning of data)"),
		until: fs.String("until", "", "end of window (empty = now)"),
	}
}

func (w windowFlags) bounds() (time.Time, time.Time, error) {
	now := time.Now()
	since, err := parseWhen(*w.since, now)
	if err != nil {
		return since, since, err
	}
	until, err := parseWhen(*w.until, now)
	return since, until, err
}

func parseWhen(s string, now time.Time) (time.Time, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return time.Time{}, nil
	}
	if d, err := time.ParseDuration(s); err == nil {
		return now.Add(-d), nil
	}
	for _, layout := range []string{time.RFC3339, "2006-01-02 15:04:05", "2006-01-02 15:04", "2006-01-02T15:04"} {
		if t, err := time.ParseInLocation(layout, s, time.Local); err == nil {
			return t, nil
		}
	}
	for _, layout := range []string{"15:04:05", "15:04"} {
		if t, err := time.ParseInLocation(layout, s, time.Local); err == nil {
			y, m, d := now.Date()
			return time.Date(y, m, d, t.Hour(), t.Minute(), t.Second(), 0, time.Local), nil
		}
	}
	return time.Time{}, fmt.Errorf("cannot parse time %q", s)
}

func cmdTop(args []string) error {
	fs := flag.NewFlagSet("top", flag.ExitOnError)
	w := addWindowFlags(fs, "1h")
	by := fs.String("by", "stall", "sort key: stall, blkio, dstate, io, read, write, cpu, majflt, rss")
	n := fs.Int("n", 20, "rows")
	fs.Parse(args)
	if sortKeys[*by] == nil {
		return fmt.Errorf("unknown sort key %q", *by)
	}
	since, until, err := w.bounds()
	if err != nil {
		return err
	}
	r := newReport()
	if err := ReadRecords(*w.dir, since, until, r.add); err != nil {
		return err
	}
	r.printSummary(os.Stdout)
	r.printProcs(os.Stdout, *by, *n)
	r.printHints(os.Stdout)
	return nil
}

func cmdBlocked(args []string) error {
	fs := flag.NewFlagSet("blocked", flag.ExitOnError)
	w := addWindowFlags(fs, "1h")
	n := fs.Int("n", 15, "rows per section")
	stacks := fs.Int("stacks", 5, "kernel stacks to show")
	fs.Parse(args)
	since, until, err := w.bounds()
	if err != nil {
		return err
	}
	r := newReport()
	if err := ReadRecords(*w.dir, since, until, r.add); err != nil {
		return err
	}
	r.printProcs(os.Stdout, "dstate", *n)
	r.printBlocked(os.Stdout, *n, *stacks)
	r.printHints(os.Stdout)
	return nil
}

func cmdIncidents(args []string) error {
	fs := flag.NewFlagSet("incidents", flag.ExitOnError)
	w := addWindowFlags(fs, "")
	n := fs.Int("n", 5, "rows per section in each incident summary")
	last := fs.Int("last", 10, "summarize only the most recent N incidents (0 = all)")
	list := fs.Bool("list", false, "list incidents without summaries")
	fs.Parse(args)
	since, until, err := w.bounds()
	if err != nil {
		return err
	}
	idx := newReport()
	if err := ReadRecords(*w.dir, since, until, idx.add); err != nil {
		return err
	}
	incs := idx.incidents
	if len(incs) == 0 {
		fmt.Println("no incidents recorded")
		return nil
	}
	fmt.Printf("%d incidents\n", len(incs))
	for _, in := range incs {
		end := "ongoing"
		if in.end != 0 {
			end = fmt.Sprintf("%s (%s)", fmtTime(in.end), time.Duration(in.end-in.start)*time.Millisecond)
		}
		fmt.Printf("  %s .. %s  %s\n", fmtTime(in.start), end, in.reason)
	}
	if *list {
		return nil
	}
	if *last > 0 && len(incs) > *last {
		incs = incs[len(incs)-*last:]
	}
	for _, in := range incs {
		end := time.UnixMilli(in.end)
		if in.end == 0 {
			end = time.Time{}
		}
		// Include a little lead-in: the cause often starts before the trigger.
		r := newReport()
		if err := ReadRecords(*w.dir, time.UnixMilli(in.start).Add(-30*time.Second), end, r.add); err != nil {
			return err
		}
		fmt.Printf("\n==== incident %s: %s ====\n", fmtTime(in.start), in.reason)
		r.printSummary(os.Stdout)
		r.printProcs(os.Stdout, "stall", *n)
		r.printBlocked(os.Stdout, *n, 1)
		r.printHints(os.Stdout)
	}
	return nil
}

func cmdShow(args []string) error {
	fs := flag.NewFlagSet("show", flag.ExitOnError)
	w := addWindowFlags(fs, "30m")
	procs := fs.Int("procs", 0, "also show the top N processes of each sample")
	blocked := fs.Bool("blocked", false, "also show each D-state scan")
	fs.Parse(args)
	since, until, err := w.bounds()
	if err != nil {
		return err
	}
	fmt.Println(sysHeader)
	maxD := 0
	return ReadRecords(*w.dir, since, until, func(rec *Record) {
		switch rec.Kind {
		case "boot":
			fmt.Printf("---- %s recorder started on %s (%s, %d CPUs)\n", fmtTime(rec.T), rec.Boot.Host, rec.Boot.Kernel, rec.Boot.NCPU)
		case "inc":
			if rec.Event == "start" {
				fmt.Printf("---- %s incident start: %s\n", fmtTime(rec.T), rec.Reason)
			} else {
				fmt.Printf("---- %s incident end after %.0fs\n", fmtTime(rec.T), rec.Dt)
			}
		case "blk":
			maxD = max(maxD, len(rec.Blocked))
			if *blocked {
				fmt.Printf("     %s D-state scan: %d threads\n", fmtTime(rec.T), len(rec.Blocked))
				printBlockedRows(os.Stdout, rec.Blocked, false)
			}
		case "sample":
			if rec.Sys != nil {
				fmt.Println(sysLine(rec.T, rec.Sys, maxD))
			}
			maxD = 0
			if *procs > 0 {
				printProcRows(os.Stdout, rec.Procs, rec.Dt, *procs, "      ")
			}
		}
	})
}

func cmdExport(args []string) error {
	fs := flag.NewFlagSet("export", flag.ExitOnError)
	w := addWindowFlags(fs, "")
	kinds := fs.String("kinds", "", "comma-separated record kinds to keep (boot,sample,blk,inc); empty = all")
	fs.Parse(args)
	since, until, err := w.bounds()
	if err != nil {
		return err
	}
	keep := map[string]bool{}
	for _, k := range strings.Split(*kinds, ",") {
		if k != "" {
			keep[k] = true
		}
	}
	enc := json.NewEncoder(os.Stdout)
	return ReadRecords(*w.dir, since, until, func(rec *Record) {
		if len(keep) == 0 || keep[rec.Kind] {
			enc.Encode(rec)
		}
	})
}
