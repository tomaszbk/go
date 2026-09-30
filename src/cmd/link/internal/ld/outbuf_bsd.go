//go:build freebsd || netbsd

package ld

import (
	"internal/syscall/unix"
	"syscall"
)

func (out *OutBuf) fallocate(size uint64) error {
	err := unix.PosixFallocate(int(out.f.Fd()), 0, int64(size))
	// ZFS on FreeBSD does not support posix_fallocate and returns EINVAL in that case.
	if err == syscall.EINVAL {
		return errNoFallocate
	}
	return err
}
