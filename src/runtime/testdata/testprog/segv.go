//go:build unix

package main

import "syscall"

func init() {
	register("Segv", Segv)
}

var Sum int

func Segv() {
	c := make(chan bool)
	go func() {
		close(c)
		for i := 0; ; i++ {
			Sum += i
		}
	}()

	<-c

	syscall.Kill(syscall.Getpid(), syscall.SIGSEGV)

	// Wait for the OS to deliver the signal.
	select {}
}
