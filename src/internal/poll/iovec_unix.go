//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd

package poll

import "syscall"

func newIovecWithBase(base *byte) syscall.Iovec {
	return syscall.Iovec{Base: base}
}
