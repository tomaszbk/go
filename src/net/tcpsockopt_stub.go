//go:build js || wasip1

package net

import (
	"syscall"
	"time"
)

func setNoDelay(fd *netFD, noDelay bool) error {
	return syscall.ENOPROTOOPT
}

func setKeepAliveIdle(fd *netFD, d time.Duration) error {
	return syscall.ENOPROTOOPT
}

func setKeepAliveInterval(fd *netFD, d time.Duration) error {
	return syscall.ENOPROTOOPT
}

func setKeepAliveCount(fd *netFD, n int) error {
	return syscall.ENOPROTOOPT
}
