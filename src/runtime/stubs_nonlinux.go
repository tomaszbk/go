//go:build !linux

package runtime

// sbrk0 returns the current process brk, or 0 if not implemented.
func sbrk0() uintptr {
	return 0
}
