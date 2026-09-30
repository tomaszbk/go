//go:build unix || windows

package net

import (
	"runtime"
	"syscall"
)

func setNoDelay(fd *netFD, noDelay bool) error {
	err := fd.pfd.SetsockoptInt(syscall.IPPROTO_TCP, syscall.TCP_NODELAY, boolint(noDelay))
	runtime.KeepAlive(fd)
	return wrapSyscallError("setsockopt", err)
}
