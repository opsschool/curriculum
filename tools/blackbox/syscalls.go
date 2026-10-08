package main

import "strconv"

type syscallInfo struct {
	name  string
	fdArg bool // first argument is a file descriptor worth resolving
}

// syscallName returns a readable name for a syscall number. Unknown numbers
// are returned as "sys_<nr>"; -1 means the thread is not in a syscall, which
// for a blocked thread almost always means it is waiting on a page fault.
func syscallName(nr int64) (string, bool) {
	if nr == -1 {
		return "pagefault", false
	}
	if s, ok := syscallTable[nr]; ok {
		return s.name, s.fdArg
	}
	return "sys_" + strconv.FormatInt(nr, 10), false
}
