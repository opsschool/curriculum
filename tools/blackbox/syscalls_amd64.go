package main

// x86_64 syscall numbers (arch/x86/entry/syscalls/syscall_64.tbl). Only used
// when running blackbox on a PC for testing.
var syscallTable = map[int64]syscallInfo{
	0: {"read", true}, 1: {"write", true}, 2: {"open", false}, 3: {"close", true},
	4: {"stat", false}, 5: {"fstat", true}, 7: {"poll", false}, 8: {"lseek", true},
	9: {"mmap", false}, 10: {"mprotect", false}, 11: {"munmap", false},
	12: {"brk", false}, 16: {"ioctl", true}, 17: {"pread64", true},
	18: {"pwrite64", true}, 19: {"readv", true}, 20: {"writev", true},
	23: {"select", false}, 24: {"sched_yield", false}, 26: {"msync", false},
	28: {"madvise", false}, 35: {"nanosleep", false}, 40: {"sendfile", true},
	56: {"clone", false}, 59: {"execve", false}, 60: {"exit", false},
	61: {"wait4", false}, 72: {"fcntl", true}, 73: {"flock", true},
	74: {"fsync", true}, 75: {"fdatasync", true}, 76: {"truncate", false},
	77: {"ftruncate", true}, 82: {"rename", false}, 83: {"mkdir", false},
	87: {"unlink", false}, 162: {"sync", false}, 202: {"futex", false},
	208: {"io_getevents", false}, 209: {"io_submit", false},
	217: {"getdents64", true}, 221: {"fadvise64", true},
	230: {"clock_nanosleep", false}, 231: {"exit_group", false},
	232: {"epoll_wait", false}, 257: {"openat", false}, 262: {"newfstatat", false},
	263: {"unlinkat", false}, 264: {"renameat", false}, 275: {"splice", true},
	277: {"sync_file_range", true}, 281: {"epoll_pwait", false},
	285: {"fallocate", true}, 295: {"preadv", true}, 296: {"pwritev", true},
	306: {"syncfs", true}, 326: {"copy_file_range", true},
	327: {"preadv2", true}, 328: {"pwritev2", true}, 332: {"statx", false},
	426: {"io_uring_enter", false},
}
