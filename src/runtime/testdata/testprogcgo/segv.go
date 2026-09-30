//go:build unix

package main

// #include <unistd.h>
// static void nop() {}
import "C"

import "syscall"

func init() {
	register("SegvInCgo", SegvInCgo)
}

func SegvInCgo() {
	c := make(chan bool)
	go func() {
		close(c)
		for {
			C.nop()
		}
	}()

	<-c

	syscall.Kill(syscall.Getpid(), syscall.SIGSEGV)

	// Wait for the OS to deliver the signal.
	C.pause()
}
