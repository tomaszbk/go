// compile


// Test escape analysis with shifting constant

package main

import "unsafe"

func main() {
	var l uint64
	var p unsafe.Pointer
	_ = unsafe.Pointer(uintptr(p) + (uintptr(l) >> 1))
}
