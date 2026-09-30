package main

import "syscall"

func init() {
	register("TgkillSegv", TgkillSegv)
}

func TgkillSegv() {
	c := make(chan bool)
	go func() {
		close(c)
		for i := 0; ; i++ {
			// Sum defined in segv.go.
			Sum += i
		}
	}()

	<-c

	syscall.Tgkill(syscall.Getpid(), syscall.Gettid(), syscall.SIGSEGV)

	// Wait for the OS to deliver the signal.
	select {}
}
