// compile


package main

import "unsafe"

func main() {
	var x int
	
	a := uint64(uintptr(unsafe.Pointer(&x)))
	b := uint32(uintptr(unsafe.Pointer(&x)))
	c := uint16(uintptr(unsafe.Pointer(&x)))
	d := int64(uintptr(unsafe.Pointer(&x)))
	e := int32(uintptr(unsafe.Pointer(&x)))
	f := int16(uintptr(unsafe.Pointer(&x)))

	_, _, _, _, _, _ = a, b, c, d, e, f
}
