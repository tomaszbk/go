//go:build unix

package unix

import (
	"runtime"
	"syscall"
)

func Eaccess(path string, mode uint32) error {
	if runtime.GOOS == "android" {
		// syscall.Faccessat for Android implements AT_EACCESS check in
		// userspace. Since Android doesn't have setuid programs and
		// never runs code with euid!=uid, AT_EACCESS check is not
		// really required. Return ENOSYS so the callers can fall back
		// to permission bits check.
		return syscall.ENOSYS
	}
	return faccessat(AT_FDCWD, path, mode, AT_EACCESS)
}
