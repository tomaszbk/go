//go:build wasip1

package net

import (
	"syscall"
)

func (fd *netFD) closeRead() error {
	if fd.fakeNetFD != nil {
		return fd.fakeNetFD.closeRead()
	}
	return fd.shutdown(syscall.SHUT_RD)
}

func (fd *netFD) closeWrite() error {
	if fd.fakeNetFD != nil {
		return fd.fakeNetFD.closeWrite()
	}
	return fd.shutdown(syscall.SHUT_WR)
}
