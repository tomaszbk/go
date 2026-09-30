//go:build !faketime

package syscall

const faketime = false

func faketimeWrite(fd int, p []byte) int {
	// This should never be called since faketime is false.
	panic("not implemented")
}
