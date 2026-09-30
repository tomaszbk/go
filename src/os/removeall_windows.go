//go:build windows

package os

import "syscall"

func newDirFile(fd syscall.Handle, name string) (*File, error) {
	return newFile(fd, name, kindOpenFile, false), nil
}
