//go:build ios

package unix

func ptrace(request int, pid int, addr uintptr, data uintptr) (err error) {
	return ENOTSUP
}
