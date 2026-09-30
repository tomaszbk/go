package os

import "syscall"

func hostname() (name string, err error) {
	return syscall.Gethostname()
}
