// errorcheck


// Test that result parameters are in the same scope as regular parameters.
// Does not compile.

package main

func f1(a int) (int, float32) {
	return 7, 7.0
}


func f2(a int) (a int, b float32) { // ERROR "duplicate argument a|definition|redeclared"
	return 8, 8.0
}
