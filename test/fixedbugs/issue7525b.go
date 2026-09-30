// errorcheck

// Issue 7525: self-referential array types.

package main

var y struct { // GC_ERROR "initialization cycle: y refers to itself"
	d [len(y.d)]int // GCCGO_ERROR "array bound|typechecking loop|invalid array"
}
