package runtime

import "unsafe"

//go:nosplit
func cputicks() int64 {
	var counter int64
	stdcall(_QueryPerformanceCounter, uintptr(unsafe.Pointer(&counter)))
	return counter
}

func stackcheck() {
	// TODO: not implemented
}
