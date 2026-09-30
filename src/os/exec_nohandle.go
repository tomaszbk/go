//go:build !linux && !windows

package os

func (ph *processHandle) closeHandle() {
	panic("internal error: unexpected call to closeHandle")
}
