//go:build !windows

package main

import "syscall"

func pipe() (r, w int, err error) {
	var p [2]int
	err = syscall.Pipe(p[:])
	return p[0], p[1], err
}
