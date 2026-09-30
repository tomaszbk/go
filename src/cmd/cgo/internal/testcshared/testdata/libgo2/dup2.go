//go:build darwin || dragonfly || freebsd || (linux && !arm64 && !loong64 && !riscv64) || netbsd || openbsd

package main

import "syscall"

func dup2(oldfd, newfd int) error {
	return syscall.Dup2(oldfd, newfd)
}
