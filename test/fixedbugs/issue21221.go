// run

package main

import "unsafe"

func main() {
	if unsafe.Pointer(uintptr(0)) != unsafe.Pointer(nil) {
		panic("fail")
	}
	if (*int)(unsafe.Pointer(uintptr(0))) != (*int)(nil) {
		panic("fail")
	}
}
