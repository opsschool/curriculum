package main

import (
	"math"
	"os"
	"runtime"
	"sort"
	"strconv"
	"time"
)

// Record is one line in the data files. Kind is one of:
//
//	boot    recorder started; Boot is set
//	sample  periodic system + process sample; Sys and Procs are set
//	blk     threads found in uninterruptible sleep (D state); Blocked is set
//	inc     incident start/end; Event and Reason are set
type Record struct {
	T    int64   `json:"t"`            // unix milliseconds
	Kind string  `json:"k"`            // see above
	Dt   float64 `json:"dt,omitempty"` // seconds covered by this record

	Boot    *Boot     `json:"boot,omitempty"`
	Sys     *Sys      `json:"sys,omitempty"`
	Procs   []Proc    `json:"procs,omitempty"`
	Full    bool      `json:"full,omitempty"` // Procs lists every process
	Blocked []Blocked `json:"blk,omitempty"`
	Load1   float64   `json:"load1,omitempty"`
	NBlk    int       `json:"nblk,omitempty"` // procs_blocked from /proc/stat
	Event   string    `json:"ev,omitempty"`
	Reason  string    `json:"reason,omitempty"`
	Dropped uint64    `json:"dropped,omitempty"` // records lost because the writer fell behind
}

type Boot struct {
	Host     string            `json:"host"`
	Kernel   string            `json:"kernel"`
	Arch     string            `json:"arch"`
	NCPU     int               `json:"ncpu"`
	Features map[string]string `json:"features"`
}

// Sys holds system-wide figures. Rates are per second, percentages are of
// total CPU capacity (all cores), memory is in kB.
type Sys struct {
	Load1   float64 `json:"l1"`
	Load5   float64 `json:"l5"`
	Load15  float64 `json:"l15"`
	Running int     `json:"run"`
	Blocked int     `json:"blk"`
	Threads int     `json:"thr"`

	User    float64 `json:"usr"`
	Nice    float64 `json:"nice,omitempty"`
	System  float64 `json:"sys"`
	IOWait  float64 `json:"iow"`
	IRQ     float64 `json:"irq,omitempty"`
	SoftIRQ float64 `json:"sirq,omitempty"`
	Steal   float64 `json:"steal,omitempty"`
	Idle    float64 `json:"idle"`
	Ctxt    float64 `json:"ctxt"`
	Forks   float64 `json:"forks"`

	MemTotal  int64 `json:"mtot"`
	MemAvail  int64 `json:"mavail"`
	MemFree   int64 `json:"mfree"`
	Buffers   int64 `json:"buf"`
	Cached    int64 `json:"cache"`
	Dirty     int64 `json:"dirty"`
	Writeback int64 `json:"wb"`
	Shmem     int64 `json:"shmem"`
	Slab      int64 `json:"slab"`
	SwapUsed  int64 `json:"swap"`

	MajFlt  float64 `json:"majflt"`
	PgIn    float64 `json:"pgin"`  // kB/s read from block devices (includes swap)
	PgOut   float64 `json:"pgout"` // kB/s written to block devices
	SwpIn   float64 `json:"swpin,omitempty"`
	SwpOut  float64 `json:"swpout,omitempty"`
	OOMKill uint64  `json:"oom,omitempty"` // OOM kills during the interval

	PSI   *PSI   `json:"psi,omitempty"`
	Disks []Disk `json:"disks,omitempty"`

	SelfRSS int64   `json:"self_rss"` // recorder's own RSS, kB
	SelfCPU float64 `json:"self_cpu"` // recorder's own CPU, % of one core
}

// PSI is the share of the interval (in %) during which some / all non-idle
// tasks were stalled on a resource. Needs CONFIG_PSI.
type PSI struct {
	CPU     float64 `json:"cpu"`
	MemSome float64 `json:"mem"`
	MemFull float64 `json:"memf"`
	IOSome  float64 `json:"io"`
	IOFull  float64 `json:"iof"`
}

type Disk struct {
	Name     string  `json:"n"`
	RIOPS    float64 `json:"r"`
	WIOPS    float64 `json:"w"`
	RKB      float64 `json:"rkb"`
	WKB      float64 `json:"wkb"`
	Util     float64 `json:"util"`  // % of time with I/O in flight
	Await    float64 `json:"await"` // ms per completed request
	InFlight int     `json:"q,omitempty"`
}

// Proc holds one process. Counters are amounts during the record's Dt, not
// rates.
type Proc struct {
	PID     int     `json:"pid"`
	PPID    int     `json:"ppid,omitempty"`
	Comm    string  `json:"comm"`
	Cmd     string  `json:"cmd,omitempty"` // on full snapshots and new processes
	State   string  `json:"st"`
	CPU     float64 `json:"cpu"` // % of one core
	RSS     int64   `json:"rss"` // kB
	Swap    int64   `json:"swap,omitempty"`
	Threads int     `json:"thr,omitempty"`
	MajFlt  int64   `json:"majflt,omitempty"`
	RdB     int64   `json:"rd,omitempty"`    // bytes read from storage
	WrB     int64   `json:"wr,omitempty"`    // bytes dirtied for writing to storage
	CWrB    int64   `json:"cwr,omitempty"`   // dirtied bytes cancelled (truncate/delete)
	RChar   int64   `json:"rchar,omitempty"` // bytes passed to read-type syscalls
	WChar   int64   `json:"wchar,omitempty"`
	BlkIO   int64   `json:"blkio,omitempty"` // ms spent waiting on block I/O (delay accounting)
}

// Blocked is one thread found in D state.
type Blocked struct {
	PID     int      `json:"pid"`
	TID     int      `json:"tid"`
	Comm    string   `json:"comm"`            // process name
	TComm   string   `json:"tcomm,omitempty"` // thread name, when different
	Wchan   string   `json:"wchan,omitempty"`
	Syscall string   `json:"sc,omitempty"`
	File    string   `json:"file,omitempty"` // target of the syscall's fd argument
	Stack   []string `json:"stack,omitempty"`
}

type Config struct {
	Interval         time.Duration
	FastInterval     time.Duration
	IncidentInterval time.Duration
	FullEvery        time.Duration
	TopN             int
	LoadThreshold    float64
	PSIThreshold     float64
	BlockedThreshold int
	IncidentHold     time.Duration
	MaxDetail        time.Duration
	MaxBlocked       int
	StackFrames      int
}

type procPrev struct {
	start  uint64
	ppid   int
	cpu    uint64
	majflt uint64
	blkio  uint64
	tblk   map[int]uint64 // per-thread blkio ticks, multi-threaded processes only
	io     pidIO
	gen    uint64
}

type diskPrev struct {
	d diskCounters
	t time.Time
}

type Collector struct {
	cfg    Config
	r      *reader
	emit   func(*Record)
	ncpu   int
	pageKB int64
	self   int

	prevT    time.Time
	prevStat statInfo
	prevVM   vmStat
	prevPSI  [3]psi
	hasPSI   bool
	disks    map[string]diskPrev
	procs    map[int]*procPrev
	gen      uint64
	lastFull time.Time
	selfCPU  uint64
	delayOn  bool

	lastProbe   time.Time
	inIncident  bool
	incStart    time.Time
	lastTrigger time.Time
	blkStreak   int
}

func NewCollector(cfg Config, emit func(*Record)) *Collector {
	c := &Collector{
		cfg:    cfg,
		r:      newReader(),
		emit:   emit,
		ncpu:   runtime.NumCPU(),
		pageKB: int64(os.Getpagesize() / 1024),
		self:   os.Getpid(),
		disks:  map[string]diskPrev{},
		procs:  map[int]*procPrev{},
	}
	c.checkDelayacct()
	if c.cfg.LoadThreshold <= 0 {
		c.cfg.LoadThreshold = float64(c.ncpu)
	}
	_, err := c.r.read("/proc/pressure/io")
	c.hasPSI = err == nil
	return c
}

// checkDelayacct notes whether per-task delay accounting is on. Kernels
// before 5.14 have no sysctl and account whenever it is compiled in.
func (c *Collector) checkDelayacct() {
	b, err := c.r.read("/proc/sys/kernel/task_delayacct")
	c.delayOn = err != nil || (len(b) > 0 && b[0] != '0')
}

func r1(x float64) float64 { return math.Round(x*10) / 10 }
func r2(x float64) float64 { return math.Round(x*100) / 100 }

// Sample reads system and process counters. The first call only primes the
// counters and returns nil.
func (c *Collector) Sample(now time.Time) *Record {
	first := c.prevT.IsZero()
	dt := now.Sub(c.prevT).Seconds()
	full := first || now.Sub(c.lastFull) >= c.cfg.FullEvery

	sys := &Sys{}
	if b, err := c.r.read("/proc/loadavg"); err == nil {
		l := parseLoadavg(b)
		sys.Load1, sys.Load5, sys.Load15, sys.Threads = l.load1, l.load5, l.load15, l.threads
	}
	if b, err := c.r.read("/proc/stat"); err == nil {
		st := parseStat(b)
		sys.Running, sys.Blocked = st.running, st.blocked
		if !first {
			p, q := st.cpu, c.prevStat.cpu
			tot := float64(sub(p.total(), q.total()))
			if tot > 0 {
				pct := func(a, b uint64) float64 { return r1(float64(sub(a, b)) * 100 / tot) }
				sys.User, sys.Nice, sys.System = pct(p.user, q.user), pct(p.nice, q.nice), pct(p.system, q.system)
				sys.IOWait, sys.IRQ, sys.SoftIRQ = pct(p.iowait, q.iowait), pct(p.irq, q.irq), pct(p.softirq, q.softirq)
				sys.Steal, sys.Idle = pct(p.steal, q.steal), pct(p.idle, q.idle)
			}
			sys.Ctxt = r1(float64(sub(st.ctxt, c.prevStat.ctxt)) / dt)
			sys.Forks = r1(float64(sub(st.forks, c.prevStat.forks)) / dt)
		}
		c.prevStat = st
	}
	if b, err := c.r.read("/proc/meminfo"); err == nil {
		m := parseMeminfo(b)
		sys.MemTotal, sys.MemAvail, sys.MemFree = m.total, m.avail, m.free
		sys.Buffers, sys.Cached, sys.Dirty, sys.Writeback = m.buffers, m.cached, m.dirty, m.writeback
		sys.Shmem, sys.Slab, sys.SwapUsed = m.shmem, m.slab, m.swapTotal-m.swapFree
	}
	if b, err := c.r.read("/proc/vmstat"); err == nil {
		v := parseVmstat(b)
		if !first {
			rate := func(a, b uint64) float64 { return r1(float64(sub(a, b)) / dt) }
			sys.MajFlt = rate(v.majflt, c.prevVM.majflt)
			sys.PgIn, sys.PgOut = rate(v.pgpgin, c.prevVM.pgpgin), rate(v.pgpgout, c.prevVM.pgpgout)
			sys.SwpIn, sys.SwpOut = rate(v.pswpin, c.prevVM.pswpin), rate(v.pswpout, c.prevVM.pswpout)
			sys.OOMKill = sub(v.oomKill, c.prevVM.oomKill)
		}
		c.prevVM = v
	}
	if c.hasPSI {
		var cur [3]psi
		for i, name := range []string{"cpu", "memory", "io"} {
			if b, err := c.r.read("/proc/pressure/" + name); err == nil {
				cur[i] = parsePSI(b)
			}
		}
		if !first {
			us := dt * 1e6
			pct := func(a, b uint64) float64 { return r1(math.Min(100, float64(sub(a, b))*100/us)) }
			sys.PSI = &PSI{
				CPU:     pct(cur[0].some.total, c.prevPSI[0].some.total),
				MemSome: pct(cur[1].some.total, c.prevPSI[1].some.total),
				MemFull: pct(cur[1].full.total, c.prevPSI[1].full.total),
				IOSome:  pct(cur[2].some.total, c.prevPSI[2].some.total),
				IOFull:  pct(cur[2].full.total, c.prevPSI[2].full.total),
			}
		}
		c.prevPSI = cur
	}
	if b, err := c.r.read("/proc/diskstats"); err == nil {
		parseDiskstats(b, func(name string, d diskCounters) {
			p, ok := c.disks[name]
			c.disks[name] = diskPrev{d, now}
			if !ok || first {
				return
			}
			q := p.d
			reads, writes := sub(d.reads, q.reads), sub(d.writes, q.writes)
			busy := sub(d.msIO, q.msIO)
			if reads == 0 && writes == 0 && busy == 0 && d.inFlight == 0 {
				return
			}
			dk := Disk{
				Name:     name,
				RIOPS:    r1(float64(reads) / dt),
				WIOPS:    r1(float64(writes) / dt),
				RKB:      r1(float64(sub(d.sectorsRead, q.sectorsRead)) / 2 / dt),
				WKB:      r1(float64(sub(d.sectorsWritten, q.sectorsWritten)) / 2 / dt),
				Util:     r1(math.Min(100, float64(busy)/(dt*10))),
				InFlight: d.inFlight,
			}
			if n := reads + writes; n > 0 {
				dk.Await = r1(float64(sub(d.msReading, q.msReading)+sub(d.msWriting, q.msWriting)) / float64(n))
			}
			sys.Disks = append(sys.Disks, dk)
		})
	}

	procs := c.collectProcs(dt, first, full)

	if b, err := c.r.read("/proc/self/stat"); err == nil {
		if ps, ok := parsePidStat(b); ok {
			sys.SelfRSS = ps.rss * c.pageKB
			if !first {
				sys.SelfCPU = r2(float64(sub(ps.cpu, c.selfCPU)) * 100 / userHZ / dt)
			}
			c.selfCPU = ps.cpu
		}
	}

	c.prevT = now
	if full {
		c.lastFull = now
		c.checkDelayacct()
	}
	if first {
		return nil
	}
	return &Record{T: now.UnixMilli(), Kind: "sample", Dt: r2(dt), Sys: sys, Procs: procs, Full: full}
}

func (c *Collector) collectProcs(dt float64, first, full bool) []Proc {
	c.gen++
	var out []Proc
	idx := map[int]int{}
	for _, pid := range listNumeric("/proc") {
		base := "/proc/" + strconv.Itoa(pid)
		b, err := c.r.read(base + "/stat")
		if err != nil {
			continue
		}
		ps, ok := parsePidStat(b)
		if !ok {
			continue
		}
		prev := c.procs[pid]
		isNew := prev == nil || prev.start != ps.start
		if isNew {
			prev = &procPrev{start: ps.start}
			c.procs[pid] = prev
		}
		io := prev.io
		if b, err := c.r.read(base + "/io"); err == nil {
			io = parsePidIO(b)
		}
		blkio := c.blkioDelta(pid, base, &ps, prev, isNew)

		p := Proc{
			PID: pid, PPID: ps.ppid, Comm: ps.comm, State: string(ps.state),
			RSS: ps.rss * c.pageKB, Threads: ps.threads,
		}
		if !first {
			// A process that appeared since the last sample started inside
			// the interval, so its counters since birth belong to it.
			p.CPU = r1(float64(sub(ps.cpu, prev.cpu)) * 100 / userHZ / dt)
			p.MajFlt = int64(sub(ps.majflt, prev.majflt))
			// Cap at what the threads could have waited; toggling
			// task_delayacct at runtime can produce one bogus jump.
			p.BlkIO = int64(math.Min(float64(blkio)*1000/userHZ, dt*1000*float64(max(ps.threads, 1))))
			p.RdB = int64(sub(io.read, prev.io.read))
			p.WrB = int64(sub(io.write, prev.io.write))
			p.CWrB = int64(sub(io.cancelled, prev.io.cancelled))
			p.RChar = int64(sub(io.rchar, prev.io.rchar))
			p.WChar = int64(sub(io.wchar, prev.io.wchar))
		}
		prev.ppid, prev.cpu, prev.majflt, prev.blkio, prev.io, prev.gen = ps.ppid, ps.cpu, ps.majflt, ps.blkio, io, c.gen

		active := p.CPU > 0 || p.MajFlt > 0 || p.BlkIO > 0 || p.RdB > 0 || p.WrB > 0 || ps.state == 'D'
		if full || (isNew && !first) {
			if b, err := c.r.read(base + "/cmdline"); err == nil {
				p.Cmd = cmdline(b, 200)
			}
		}
		if full {
			if b, err := c.r.read(base + "/status"); err == nil {
				p.Swap = vmSwap(b)
			}
			idx[pid] = len(out)
			out = append(out, p)
		} else if active || (isNew && !first) {
			idx[pid] = len(out)
			out = append(out, p)
		}
	}
	for pid, pp := range c.procs {
		if pp.gen == c.gen {
			continue
		}
		// When a parent reaps a child, the kernel adds the child's I/O
		// counters to the parent's. Take back what was already counted for
		// the child so the bytes are not reported twice.
		if i, ok := idx[pp.ppid]; ok && !first {
			p := &out[i]
			p.RdB = max(0, p.RdB-int64(pp.io.read))
			p.WrB = max(0, p.WrB-int64(pp.io.write))
			p.CWrB = max(0, p.CWrB-int64(pp.io.cancelled))
			p.RChar = max(0, p.RChar-int64(pp.io.rchar))
			p.WChar = max(0, p.WChar-int64(pp.io.wchar))
		}
		delete(c.procs, pid)
	}
	if !full && len(out) > c.cfg.TopN {
		sort.Slice(out, func(i, j int) bool { return procScore(&out[i]) > procScore(&out[j]) })
		out = out[:c.cfg.TopN]
	}
	return out
}

// blkioDelta returns the block I/O delay ticks a process accumulated since
// the last sample. /proc/<pid>/stat reports the delay of the main thread
// only, so multi-threaded processes are summed over their threads.
func (c *Collector) blkioDelta(pid int, base string, ps *pidStat, prev *procPrev, isNew bool) uint64 {
	if !c.delayOn {
		return 0
	}
	if ps.threads <= 1 {
		prev.tblk = nil
		return sub(ps.blkio, prev.blkio)
	}
	old := prev.tblk
	if old == nil && !isNew {
		// Was single-threaded last time: the main thread's value is known.
		old = map[int]uint64{pid: prev.blkio}
	}
	next := make(map[int]uint64, ps.threads)
	var d uint64
	for _, tid := range listNumeric(base + "/task") {
		b, err := c.r.read(base + "/task/" + strconv.Itoa(tid) + "/stat")
		if err != nil {
			continue
		}
		t, ok := parsePidStat(b)
		if !ok {
			continue
		}
		next[tid] = t.blkio
		d += sub(t.blkio, old[tid]) // a thread new since last time started at 0
	}
	prev.tblk = next
	return d
}

// procScore ranks processes for the per-sample top-N cut.
func procScore(p *Proc) float64 {
	s := p.CPU + float64(p.BlkIO)/10 + float64(p.RdB+p.WrB)/(256<<10) + float64(p.MajFlt)/10
	if p.State == "D" {
		s += 50
	}
	return s
}

// Probe runs on the fast tick. It decides whether an incident is in progress
// and, when anything is blocked, records which threads are in D state.
func (c *Collector) Probe(now time.Time) (inIncident bool) {
	dt := c.cfg.FastInterval.Seconds()
	if !c.lastProbe.IsZero() {
		dt = now.Sub(c.lastProbe).Seconds()
	}
	c.lastProbe = now

	var load1 float64
	if b, err := c.r.read("/proc/loadavg"); err == nil {
		load1 = parseLoadavg(b).load1
	}
	blocked := 0
	if b, err := c.r.read("/proc/stat"); err == nil {
		blocked = parseStat(b).blocked
	}
	var ioSome float64
	if c.hasPSI {
		if b, err := c.r.read("/proc/pressure/io"); err == nil {
			ioSome = parsePSI(b).some.avg10
		}
	}

	if blocked >= c.cfg.BlockedThreshold {
		c.blkStreak++
	} else {
		c.blkStreak = 0
	}
	reason := ""
	switch {
	case load1 >= c.cfg.LoadThreshold:
		reason = "load1=" + strconv.FormatFloat(load1, 'f', 2, 64) + " >= " + strconv.FormatFloat(c.cfg.LoadThreshold, 'f', 1, 64)
	case c.hasPSI && c.cfg.PSIThreshold > 0 && ioSome >= c.cfg.PSIThreshold:
		reason = "psi io avg10=" + strconv.FormatFloat(ioSome, 'f', 1, 64) + "%"
	case c.blkStreak >= 3:
		reason = "procs_blocked=" + strconv.Itoa(blocked) + " for 3 probes"
	}
	if reason != "" {
		c.lastTrigger = now
		if !c.inIncident {
			c.inIncident, c.incStart = true, now
			c.emit(&Record{T: now.UnixMilli(), Kind: "inc", Event: "start", Reason: reason, Load1: load1, NBlk: blocked})
		}
	} else if c.inIncident && now.Sub(c.lastTrigger) >= c.cfg.IncidentHold {
		c.inIncident = false
		c.emit(&Record{T: now.UnixMilli(), Kind: "inc", Event: "end", Dt: r1(now.Sub(c.incStart).Seconds())})
	}

	detail := c.inIncident && now.Sub(c.incStart) < c.cfg.MaxDetail
	if blocked > 0 || detail {
		if blk := c.ScanBlocked(detail); len(blk) > 0 {
			c.emit(&Record{T: now.UnixMilli(), Kind: "blk", Dt: r2(dt), Blocked: blk, Load1: load1, NBlk: blocked})
		}
	}
	return c.inIncident
}

// Stop closes an open incident.
func (c *Collector) Stop(now time.Time) {
	if c.inIncident {
		c.inIncident = false
		c.emit(&Record{T: now.UnixMilli(), Kind: "inc", Event: "end", Dt: r1(now.Sub(c.incStart).Seconds()), Reason: "recorder stopped"})
	}
}

// Detailed reports whether probes are currently in detailed (fast) mode.
func (c *Collector) Detailed(now time.Time) bool {
	return c.inIncident && now.Sub(c.incStart) < c.cfg.MaxDetail
}

// ScanBlocked finds every thread in D state. With detail it also captures
// the kernel stack of each one.
func (c *Collector) ScanBlocked(detail bool) []Blocked {
	var out []Blocked
	for _, pid := range listNumeric("/proc") {
		base := "/proc/" + strconv.Itoa(pid)
		b, err := c.r.read(base + "/stat")
		if err != nil {
			continue
		}
		ps, ok := parsePidStat(b)
		if !ok {
			continue
		}
		// Single-threaded processes (and kernel threads) need no task walk.
		var tids []int
		if ps.threads <= 1 {
			if ps.state != 'D' {
				continue
			}
			tids = []int{pid}
		} else {
			tids = listNumeric(base + "/task")
		}
		for _, tid := range tids {
			tbase := base + "/task/" + strconv.Itoa(tid)
			b, err := c.r.read(tbase + "/stat")
			if err != nil || statState(b) != 'D' {
				continue
			}
			bl := Blocked{PID: pid, TID: tid, Comm: ps.comm}
			if tc := statComm(b); tc != ps.comm {
				bl.TComm = tc
			}
			if b, err := c.r.read(tbase + "/wchan"); err == nil && len(b) > 0 && b[0] != '0' {
				bl.Wchan = string(b)
			}
			// Kernel threads (children of kthreadd) have no user syscall.
			kthread := pid == 2 || ps.ppid == 2
			if b, err := c.r.read(tbase + "/syscall"); err == nil && !kthread {
				nr, arg0, running := parseSyscall(b)
				if running {
					bl.Syscall = "running" // woke up between the two reads
				} else {
					name, fdArg := syscallName(nr)
					bl.Syscall = name
					if fdArg && arg0 < 1<<20 {
						if target, err := os.Readlink(base + "/fd/" + strconv.FormatUint(arg0, 10)); err == nil {
							bl.File = target
						}
					}
				}
			}
			if detail {
				if b, err := c.r.read(tbase + "/stack"); err == nil {
					bl.Stack = parseStack(b, c.cfg.StackFrames)
				}
			}
			out = append(out, bl)
			if len(out) >= c.cfg.MaxBlocked {
				return out
			}
		}
	}
	return out
}
