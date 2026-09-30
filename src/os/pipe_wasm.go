//go:build wasm

package os

import "syscall"

// Pipe returns a connected pair of Files; reads from r return bytes written to w.
// It returns the files and an error, if any.
func Pipe() (r *File, w *File, err error) {
	// Neither GOOS=js nor GOOS=wasip1 have pipes.
	return nil, nil, NewSyscallError("pipe", syscall.ENOSYS)
}
