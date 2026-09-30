//go:build !go1.21 && (aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || zos)

package unix

import "syscall"

func Auxv() ([][2]uintptr, error) {
	return nil, syscall.ENOTSUP
}
