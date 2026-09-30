// errorcheck


package main

func f (x int) {	// GCCGO_ERROR "previous"
	var x int;	// ERROR "redecl|redefinition"
}
