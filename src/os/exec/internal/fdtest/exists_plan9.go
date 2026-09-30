//go:build plan9

package fdtest

import (
	"syscall"
)

const errBadFd = syscall.ErrorString("fd out of range or not open")

// Exists returns true if fd is a valid file descriptor.
func Exists(fd uintptr) bool {
	var buf [1]byte
	_, err := syscall.Fstat(int(fd), buf[:])
	return err != errBadFd
}
