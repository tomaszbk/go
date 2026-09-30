// run

//go:build !plan9 && !windows && !wasip1


// Test that a program can survive SIGCHLD.

package main

import "syscall"

func main() {
	syscall.Kill(syscall.Getpid(), syscall.SIGCHLD)
	println("survived SIGCHLD")
}
