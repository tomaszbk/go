package net

import (
	"syscall"
	"time"
)

func setKeepAliveIdle(_ *netFD, d time.Duration) error {
	if d < 0 {
		return nil
	}
	// OpenBSD has no user-settable per-socket TCP keepalive
	// options.
	return syscall.ENOPROTOOPT
}

func setKeepAliveInterval(_ *netFD, d time.Duration) error {
	if d < 0 {
		return nil
	}
	// OpenBSD has no user-settable per-socket TCP keepalive
	// options.
	return syscall.ENOPROTOOPT
}

func setKeepAliveCount(_ *netFD, n int) error {
	if n < 0 {
		return nil
	}
	// OpenBSD has no user-settable per-socket TCP keepalive
	// options.
	return syscall.ENOPROTOOPT
}
