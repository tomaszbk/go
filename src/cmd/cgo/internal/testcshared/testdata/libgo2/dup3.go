//go:build (linux && arm64) || (linux && loong64) || (linux && riscv64)

package main

import "syscall"

func dup2(oldfd, newfd int) error {
	return syscall.Dup3(oldfd, newfd, 0)
}
