// errorcheck

// Issue 7525: self-referential array types.

package main

import "unsafe"

var x struct { // GC_ERROR "initialization cycle: x refers to itself"
	b [unsafe.Offsetof(x.b)]int // GCCGO_ERROR "array bound|typechecking loop|invalid array"
}
