//go:build wasip1

package unix

import "syscall"

func Fcntl(fd int, cmd int, arg int) (int, error) {
	if cmd == syscall.F_GETFL {
		flags, err := fd_fdstat_get_flags(fd)
		return int(flags), err
	}
	return 0, syscall.ENOSYS
}
