//go:build js && wasm

package unix

import "syscall"

func Fcntl(fd int, cmd int, arg int) (int, error) {
	return 0, syscall.ENOSYS
}
