//go:build unix

package poll

import "syscall"

// SetsockoptByte wraps the setsockopt network call with a byte argument.
func (fd *FD) SetsockoptByte(level, name int, arg byte) error {
	if err := fd.incref(); err != nil {
		return err
	}
	defer fd.decref()
	return syscall.SetsockoptByte(fd.Sysfd, level, name, arg)
}
