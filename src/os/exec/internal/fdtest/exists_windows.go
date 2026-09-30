//go:build windows

package fdtest

// Exists is not implemented on windows and panics.
func Exists(fd uintptr) bool {
	panic("unimplemented")
}
