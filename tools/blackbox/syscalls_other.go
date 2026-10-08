//go:build !arm64 && !amd64

package main

// No name table for this architecture; syscalls are reported by number.
var syscallTable = map[int64]syscallInfo{}
