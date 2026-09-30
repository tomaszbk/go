package main

import "syscall"

func pipe() (r, w syscall.Handle, err error) {
	var p [2]syscall.Handle
	err = syscall.Pipe(p[:])
	return p[0], p[1], err
}
