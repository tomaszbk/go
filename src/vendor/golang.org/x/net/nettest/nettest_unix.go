//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || zos

package nettest

import "syscall"

func supportsRawSocket() bool {
	for _, af := range []int{syscall.AF_INET, syscall.AF_INET6} {
		s, err := syscall.Socket(af, syscall.SOCK_RAW, 0)
		if err != nil {
			continue
		}
		syscall.Close(s)
		return true
	}
	return false
}
