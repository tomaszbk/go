package ld

import "syscall"

func (out *OutBuf) fallocate(size uint64) error {
	return syscall.Fallocate(int(out.f.Fd()), 0, 0, int64(size))
}
