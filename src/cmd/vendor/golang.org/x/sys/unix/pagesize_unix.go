//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || zos

// For Unix, get the pagesize from the runtime.

package unix

import "syscall"

func Getpagesize() int {
	return syscall.Getpagesize()
}
