//go:build darwin && !ios

package unix

func ptrace(request int, pid int, addr uintptr, data uintptr) error {
	return ptrace1(request, pid, addr, data)
}
