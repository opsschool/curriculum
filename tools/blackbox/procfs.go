package main

// Small, allocation-conscious readers for the /proc files blackbox uses.
// Everything here works on a plain procfs; nothing needs eBPF, perf, or
// kernel modules.

import (
	"bytes"
	"os"
	"sort"
	"strconv"
	"strings"
	"syscall"
)

// userHZ is the unit of the tick counters in /proc/stat and /proc/<pid>/stat.
// It is 100 on every mainstream Linux architecture, including arm and arm64.
const userHZ = 100

// reader reads small files into a reused buffer. The returned slice is only
// valid until the next call.
type reader struct{ buf []byte }

func newReader() *reader { return &reader{buf: make([]byte, 16<<10)} }

func (r *reader) read(path string) ([]byte, error) {
	fd, err := syscall.Open(path, syscall.O_RDONLY|syscall.O_CLOEXEC, 0)
	if err != nil {
		return nil, err
	}
	defer syscall.Close(fd)
	n := 0
	for {
		if n == len(r.buf) {
			r.buf = append(r.buf, make([]byte, len(r.buf))...)
		}
		m, err := syscall.Read(fd, r.buf[n:])
		if err == syscall.EINTR {
			continue
		}
		if err != nil {
			return nil, err
		}
		if m == 0 {
			break
		}
		n += m
	}
	return r.buf[:n], nil
}

func nextField(b []byte) (field, rest []byte) {
	i := 0
	for i < len(b) && (b[i] == ' ' || b[i] == '\t' || b[i] == '\n') {
		i++
	}
	j := i
	for j < len(b) && b[j] != ' ' && b[j] != '\t' && b[j] != '\n' {
		j++
	}
	return b[i:j], b[j:]
}

func atou(b []byte) uint64 {
	var n uint64
	for _, c := range b {
		if c < '0' || c > '9' {
			break
		}
		n = n*10 + uint64(c-'0')
	}
	return n
}

func atoi(b []byte) int64 {
	if len(b) > 0 && b[0] == '-' {
		return -int64(atou(b[1:]))
	}
	return int64(atou(b))
}

func atof(b []byte) float64 {
	f, _ := strconv.ParseFloat(string(b), 64)
	return f
}

func eachLine(b []byte, fn func(line []byte)) {
	for len(b) > 0 {
		i := bytes.IndexByte(b, '\n')
		if i < 0 {
			fn(b)
			return
		}
		fn(b[:i])
		b = b[i+1:]
	}
}

// sub returns a-b, or 0 when a counter went backwards (wrap, pid reuse).
func sub(a, b uint64) uint64 {
	if a < b {
		return 0
	}
	return a - b
}

// listNumeric returns the numeric entries of a /proc-style directory.
func listNumeric(dir string) []int {
	f, err := os.Open(dir)
	if err != nil {
		return nil
	}
	defer f.Close()
	names, _ := f.Readdirnames(-1)
	out := make([]int, 0, len(names))
	for _, n := range names {
		if n == "" || n[0] < '0' || n[0] > '9' {
			continue
		}
		if v, err := strconv.Atoi(n); err == nil {
			out = append(out, v)
		}
	}
	sort.Ints(out)
	return out
}

// ---- /proc/stat ----

type cpuTimes struct {
	user, nice, system, idle, iowait, irq, softirq, steal uint64
}

func (c cpuTimes) total() uint64 {
	return c.user + c.nice + c.system + c.idle + c.iowait + c.irq + c.softirq + c.steal
}

type statInfo struct {
	cpu              cpuTimes
	ctxt, forks      uint64
	running, blocked int
}

var (
	pfxCPU     = []byte("cpu ")
	pfxCtxt    = []byte("ctxt ")
	pfxProcs   = []byte("processes ")
	pfxRunning = []byte("procs_running ")
	pfxBlocked = []byte("procs_blocked ")
)

func parseStat(b []byte) (s statInfo) {
	eachLine(b, func(l []byte) {
		switch {
		case bytes.HasPrefix(l, pfxCPU):
			f := l[len(pfxCPU):]
			var v [8]uint64
			for i := range v {
				var x []byte
				x, f = nextField(f)
				v[i] = atou(x)
			}
			s.cpu = cpuTimes{v[0], v[1], v[2], v[3], v[4], v[5], v[6], v[7]}
		case bytes.HasPrefix(l, pfxCtxt):
			s.ctxt = atou(l[len(pfxCtxt):])
		case bytes.HasPrefix(l, pfxProcs):
			s.forks = atou(l[len(pfxProcs):])
		case bytes.HasPrefix(l, pfxRunning):
			s.running = int(atou(l[len(pfxRunning):]))
		case bytes.HasPrefix(l, pfxBlocked):
			s.blocked = int(atou(l[len(pfxBlocked):]))
		}
	})
	return s
}

// ---- /proc/loadavg ----

type loadAvg struct {
	load1, load5, load15 float64
	threads              int
}

func parseLoadavg(b []byte) (l loadAvg) {
	var f []byte
	f, b = nextField(b)
	l.load1 = atof(f)
	f, b = nextField(b)
	l.load5 = atof(f)
	f, b = nextField(b)
	l.load15 = atof(f)
	f, _ = nextField(b)
	if i := bytes.IndexByte(f, '/'); i >= 0 {
		l.threads = int(atou(f[i+1:]))
	}
	return l
}

// ---- /proc/meminfo (values in kB) ----

type memInfo struct {
	total, free, avail, buffers, cached, dirty, writeback, shmem, slab, swapTotal, swapFree int64
}

func parseMeminfo(b []byte) (m memInfo) {
	eachLine(b, func(l []byte) {
		i := bytes.IndexByte(l, ':')
		if i < 0 {
			return
		}
		v, _ := nextField(l[i+1:])
		n := int64(atou(v))
		switch string(l[:i]) {
		case "MemTotal":
			m.total = n
		case "MemFree":
			m.free = n
		case "MemAvailable":
			m.avail = n
		case "Buffers":
			m.buffers = n
		case "Cached":
			m.cached = n
		case "Dirty":
			m.dirty = n
		case "Writeback":
			m.writeback = n
		case "Shmem":
			m.shmem = n
		case "Slab":
			m.slab = n
		case "SwapTotal":
			m.swapTotal = n
		case "SwapFree":
			m.swapFree = n
		}
	})
	return m
}

// ---- /proc/vmstat ----

type vmStat struct {
	majflt, pgpgin, pgpgout, pswpin, pswpout, oomKill uint64
}

func parseVmstat(b []byte) (v vmStat) {
	eachLine(b, func(l []byte) {
		k, rest := nextField(l)
		val, _ := nextField(rest)
		switch string(k) {
		case "pgmajfault":
			v.majflt = atou(val)
		case "pgpgin":
			v.pgpgin = atou(val)
		case "pgpgout":
			v.pgpgout = atou(val)
		case "pswpin":
			v.pswpin = atou(val)
		case "pswpout":
			v.pswpout = atou(val)
		case "oom_kill":
			v.oomKill = atou(val)
		}
	})
	return v
}

// ---- /proc/pressure/{cpu,memory,io} ----

type psiLine struct {
	avg10 float64
	total uint64 // microseconds stalled, cumulative
}

type psi struct{ some, full psiLine }

func parsePSI(b []byte) (p psi) {
	eachLine(b, func(l []byte) {
		kind, rest := nextField(l)
		var dst *psiLine
		switch string(kind) {
		case "some":
			dst = &p.some
		case "full":
			dst = &p.full
		default:
			return
		}
		for {
			var f []byte
			f, rest = nextField(rest)
			if len(f) == 0 {
				break
			}
			if bytes.HasPrefix(f, []byte("avg10=")) {
				dst.avg10 = atof(f[6:])
			} else if bytes.HasPrefix(f, []byte("total=")) {
				dst.total = atou(f[6:])
			}
		}
	})
	return p
}

// ---- /proc/diskstats ----

type diskCounters struct {
	reads, sectorsRead, msReading     uint64
	writes, sectorsWritten, msWriting uint64
	inFlight                          int
	msIO                              uint64
}

// parseDiskstats calls fn for every device that is not a loop or ram disk.
func parseDiskstats(b []byte, fn func(name string, d diskCounters)) {
	eachLine(b, func(l []byte) {
		var f []byte
		_, l = nextField(l) // major
		_, l = nextField(l) // minor
		name, l := nextField(l)
		if len(name) == 0 || bytes.HasPrefix(name, []byte("loop")) || bytes.HasPrefix(name, []byte("ram")) {
			return
		}
		var v [10]uint64
		for i := range v {
			f, l = nextField(l)
			v[i] = atou(f)
		}
		fn(string(name), diskCounters{
			reads: v[0], sectorsRead: v[2], msReading: v[3],
			writes: v[4], sectorsWritten: v[6], msWriting: v[7],
			inFlight: int(v[8]), msIO: v[9],
		})
	})
}

// ---- /proc/<pid>/stat ----

type pidStat struct {
	comm    string
	state   byte
	ppid    int
	majflt  uint64
	cpu     uint64 // utime+stime, ticks
	threads int
	start   uint64 // ticks after boot
	rss     int64  // pages
	blkio   uint64 // delayacct_blkio_ticks
}

// statState returns the state letter from a /proc/<pid>/stat line without
// parsing anything else.
func statState(b []byte) byte {
	i := bytes.LastIndexByte(b, ')')
	if i < 0 || i+2 >= len(b) {
		return 0
	}
	return b[i+2]
}

func statComm(b []byte) string {
	i := bytes.IndexByte(b, '(')
	j := bytes.LastIndexByte(b, ')')
	if i < 0 || j < i {
		return ""
	}
	return string(b[i+1 : j])
}

func parsePidStat(b []byte) (p pidStat, ok bool) {
	open := bytes.IndexByte(b, '(')
	close := bytes.LastIndexByte(b, ')')
	if open < 0 || close < open || close+2 >= len(b) {
		return p, false
	}
	p.comm = string(b[open+1 : close])
	f := b[close+2:]
	var utime uint64
	for i := 0; ; i++ {
		var fld []byte
		fld, f = nextField(f)
		if len(fld) == 0 {
			break
		}
		// Index 0 is field 3 (state) in proc(5) numbering.
		switch i {
		case 0:
			p.state = fld[0]
		case 1:
			p.ppid = int(atoi(fld))
		case 9:
			p.majflt = atou(fld)
		case 11:
			utime = atou(fld)
		case 12:
			p.cpu = utime + atou(fld)
		case 17:
			p.threads = int(atoi(fld))
		case 19:
			p.start = atou(fld)
		case 21:
			p.rss = atoi(fld)
		case 39:
			p.blkio = atou(fld)
			return p, true
		}
	}
	return p, true
}

// ---- /proc/<pid>/io ----

type pidIO struct {
	rchar, wchar, read, write, cancelled uint64
}

func parsePidIO(b []byte) (io pidIO) {
	eachLine(b, func(l []byte) {
		i := bytes.IndexByte(l, ':')
		if i < 0 {
			return
		}
		v, _ := nextField(l[i+1:])
		n := atou(v)
		switch string(l[:i]) {
		case "rchar":
			io.rchar = n
		case "wchar":
			io.wchar = n
		case "read_bytes":
			io.read = n
		case "write_bytes":
			io.write = n
		case "cancelled_write_bytes":
			io.cancelled = n
		}
	})
	return io
}

// vmSwap returns VmSwap (kB) from /proc/<pid>/status.
func vmSwap(b []byte) int64 {
	i := bytes.Index(b, []byte("\nVmSwap:"))
	if i < 0 {
		return 0
	}
	v, _ := nextField(b[i+8:])
	return int64(atou(v))
}

// cmdline turns a NUL-separated /proc/<pid>/cmdline into a printable string.
func cmdline(b []byte, max int) string {
	b = bytes.TrimRight(b, "\x00")
	if len(b) > max {
		b = b[:max]
	}
	s := strings.Map(func(r rune) rune {
		if r < ' ' {
			return ' '
		}
		return r
	}, string(b))
	return s
}

// ---- /proc/<pid>/task/<tid>/{syscall,stack,wchan} ----

// parseSyscall returns the syscall number and first argument. nr is -1 when
// the thread is blocked outside a syscall (typically in a page fault), and
// running is true when the thread was on CPU.
func parseSyscall(b []byte) (nr int64, arg0 uint64, running bool) {
	f, rest := nextField(b)
	if string(f) == "running" {
		return 0, 0, true
	}
	nr = atoi(f)
	a, _ := nextField(rest)
	a = bytes.TrimPrefix(a, []byte("0x"))
	arg0, _ = strconv.ParseUint(string(a), 16, 64)
	return nr, arg0, false
}

// parseStack turns /proc/<pid>/stack output into bare function names:
// "[<0>] io_schedule+0x20/0x50" becomes "io_schedule".
func parseStack(b []byte, maxFrames int) []string {
	var out []string
	eachLine(b, func(l []byte) {
		if len(out) >= maxFrames {
			return
		}
		if i := bytes.Index(l, []byte("] ")); i >= 0 {
			l = l[i+2:]
		}
		if i := bytes.IndexByte(l, '+'); i >= 0 {
			l = l[:i]
		}
		if len(l) > 0 {
			out = append(out, string(l))
		}
	})
	return out
}
