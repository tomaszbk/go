//go:build unix || (js && wasm)

package os

import (
	"internal/poll"
	"syscall"
)

func open(path string, flag int, perm uint32) (int, poll.SysFile, error) {
	fd, err := syscall.Open(path, flag, perm)
	return fd, poll.SysFile{}, err
}
