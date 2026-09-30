// run


package main

import "unsafe"

func main() {
	// works
	addr := uintptr(0x234)
	x1 := (*int)(unsafe.Pointer(addr))

	// fails
	x2 := (*int)(unsafe.Pointer(uintptr(0x234)))

	if x1 != x2 {
		println("mismatch", x1, x2)
		panic("fail")
	}
}
