//go:build unix || wasip1

package poll

import (
	"internal/syscall/unix"
	"syscall"
)

func (fd *FD) Fstatat(name string, s *syscall.Stat_t, flags int) error {
	if err := fd.incref(); err != nil {
		return err
	}
	defer fd.decref()
	return ignoringEINTR(func() error {
		return unix.Fstatat(fd.Sysfd, name, s, flags)
	})
}
