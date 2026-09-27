//go:build !windows

package main

import (
	"os"
	"syscall"
)

func execProgram(program string, args []string) error {
	return syscall.Exec(program, append([]string{program}, args...), os.Environ())
}
