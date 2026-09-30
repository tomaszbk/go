// build

// Test that the compiler does not crash during compilation.

package main

import "unsafe"

// Issue 7413
func f1() {
	type t struct {
		i int
	}

	var v *t
	_ = int(uintptr(unsafe.Pointer(&v.i)))
	_ = int32(uintptr(unsafe.Pointer(&v.i)))
}

func main() {}
