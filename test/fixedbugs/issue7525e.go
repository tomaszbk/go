// errorcheck

// Issue 7525: self-referential array types.

package main

import "unsafe"

var x struct { // GC_ERROR "initialization cycle: x refers to itself"
	c [unsafe.Alignof(x.c)]int // GCCGO_ERROR "array bound|typechecking loop|invalid array"
}
