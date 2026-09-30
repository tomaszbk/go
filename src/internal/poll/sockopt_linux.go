package poll

import "syscall"

// SetsockoptIPMreqn wraps the setsockopt network call with an IPMreqn argument.
func (fd *FD) SetsockoptIPMreqn(level, name int, mreq *syscall.IPMreqn) error {
	if err := fd.incref(); err != nil {
		return err
	}
	defer fd.decref()
	return syscall.SetsockoptIPMreqn(fd.Sysfd, level, name, mreq)
}
