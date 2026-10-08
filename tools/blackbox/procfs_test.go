package main

import (
	"reflect"
	"testing"
)

func TestParsePidStat(t *testing.T) {
	// comm containing spaces and parentheses must not shift the fields.
	line := "1234 (my (odd) name) D 1 1234 1234 0 -1 4194560 100 0 7 0 250 50 0 0 20 0 3 0 5555 10000000 640 " +
		"18446744073709551615 1 1 0 0 0 0 0 0 0 0 0 0 17 2 0 0 42 0 0 0 0 0 0 0 0 0 0\n"
	p, ok := parsePidStat([]byte(line))
	if !ok {
		t.Fatal("parse failed")
	}
	want := pidStat{comm: "my (odd) name", state: 'D', ppid: 1, majflt: 7, cpu: 300, threads: 3, start: 5555, rss: 640, blkio: 42}
	if p != want {
		t.Fatalf("got %+v\nwant %+v", p, want)
	}
	if statState([]byte(line)) != 'D' || statComm([]byte(line)) != "my (odd) name" {
		t.Fatal("statState/statComm")
	}
}

func TestParseStat(t *testing.T) {
	b := []byte("cpu  10 1 20 300 40 2 3 0 0 0\ncpu0 1 1 1 1 1 1 1 0 0 0\nintr 1 2 3\nctxt 999\nbtime 1\nprocesses 77\nprocs_running 3\nprocs_blocked 5\n")
	s := parseStat(b)
	if s.cpu != (cpuTimes{10, 1, 20, 300, 40, 2, 3, 0}) || s.ctxt != 999 || s.forks != 77 || s.running != 3 || s.blocked != 5 {
		t.Fatalf("%+v", s)
	}
}

func TestParsePSI(t *testing.T) {
	b := []byte("some avg10=12.50 avg60=3.00 avg300=1.00 total=123456\nfull avg10=4.25 avg60=1.00 avg300=0.50 total=6543\n")
	p := parsePSI(b)
	if p.some.avg10 != 12.5 || p.some.total != 123456 || p.full.avg10 != 4.25 || p.full.total != 6543 {
		t.Fatalf("%+v", p)
	}
}

func TestParseDiskstats(t *testing.T) {
	b := []byte("   7       0 loop0 1 0 2 0 0 0 0 0 0 0 0\n 179       0 mmcblk0 100 5 800 40 200 7 1600 90 2 120 130 0 0 0 0\n")
	var got []string
	parseDiskstats(b, func(name string, d diskCounters) {
		got = append(got, name)
		want := diskCounters{reads: 100, sectorsRead: 800, msReading: 40, writes: 200, sectorsWritten: 1600, msWriting: 90, inFlight: 2, msIO: 120}
		if d != want {
			t.Fatalf("%+v", d)
		}
	})
	if !reflect.DeepEqual(got, []string{"mmcblk0"}) {
		t.Fatal(got)
	}
}

func TestParseSyscallAndStack(t *testing.T) {
	nr, arg0, running := parseSyscall([]byte("82 0x7 0xffff 0x0 0x0 0x0 0x0 0xfff 0xaaa\n"))
	if nr != 82 || arg0 != 7 || running {
		t.Fatal(nr, arg0, running)
	}
	if nr, _, _ := parseSyscall([]byte("-1 0xffff 0xaaaa\n")); nr != -1 {
		t.Fatal(nr)
	}
	if _, _, running := parseSyscall([]byte("running\n")); !running {
		t.Fatal("running")
	}
	st := parseStack([]byte("[<0>] jbd2_log_wait_commit+0xb8/0x130\n[<0>] ext4_sync_file+0x1f0/0x330\n"), 1)
	if !reflect.DeepEqual(st, []string{"jbd2_log_wait_commit"}) {
		t.Fatal(st)
	}
}

func TestParseMeminfoAndIO(t *testing.T) {
	m := parseMeminfo([]byte("MemTotal:  1000 kB\nMemAvailable:  400 kB\nDirty:  12 kB\nSwapTotal: 100 kB\nSwapFree: 60 kB\n"))
	if m.total != 1000 || m.avail != 400 || m.dirty != 12 || m.swapTotal-m.swapFree != 40 {
		t.Fatalf("%+v", m)
	}
	io := parsePidIO([]byte("rchar: 5\nwchar: 6\nsyscr: 1\nsyscw: 1\nread_bytes: 4096\nwrite_bytes: 8192\ncancelled_write_bytes: 0\n"))
	if io != (pidIO{rchar: 5, wchar: 6, read: 4096, write: 8192}) {
		t.Fatalf("%+v", io)
	}
}
