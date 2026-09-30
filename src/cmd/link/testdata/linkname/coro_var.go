// Linkname "var" to reference newcoro is not allowed.

package main

import "unsafe"

func main() {
	call(&newcoro)
}

//go:linkname newcoro runtime.newcoro
var newcoro unsafe.Pointer

//go:noinline
func call(*unsafe.Pointer) {
	// not implemented
}
