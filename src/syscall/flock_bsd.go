//go:build darwin || dragonfly || freebsd || netbsd || openbsd

package syscall

import "unsafe"

// FcntlFlock performs a fcntl syscall for the [F_GETLK], [F_SETLK] or [F_SETLKW] command.
func FcntlFlock(fd uintptr, cmd int, lk *Flock_t) error {
	_, err := fcntlPtr(int(fd), cmd, unsafe.Pointer(lk))
	return err
}
