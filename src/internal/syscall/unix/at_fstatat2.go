//go:build freebsd || (linux && (loong64 || mips64 || mips64le))

package unix

import "syscall"

func Fstatat(dirfd int, path string, stat *syscall.Stat_t, flags int) error {
	return syscall.Fstatat(dirfd, path, stat, flags)
}
