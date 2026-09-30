package syscall

import "unsafe"

// On AIX, there is no flock() system call.

// FcntlFlock performs a fcntl syscall for the [F_GETLK], [F_SETLK] or [F_SETLKW] command.
func FcntlFlock(fd uintptr, cmd int, lk *Flock_t) (err error) {
	_, _, e1 := syscall6(uintptr(unsafe.Pointer(&libc_fcntl)), 3, uintptr(fd), uintptr(cmd), uintptr(unsafe.Pointer(lk)), 0, 0, 0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
