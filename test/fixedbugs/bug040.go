// errorcheck


package main

func f (x,		// GCCGO_ERROR "previous"
	x int) {	// ERROR "duplicate argument|redefinition|redeclared"
}
