//go:build !darwin && !freebsd && !linux && !netbsd

package ld

func (out *OutBuf) fallocate(size uint64) error {
	return errNoFallocate
}
