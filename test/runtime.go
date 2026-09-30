// errorcheck


// Test that even if a file imports runtime,
// it cannot get at the low-level runtime definitions
// known to the compiler.  For normal packages
// the compiler doesn't even record the lower case
// functions in its symbol table, but some functions
// in runtime are hard-coded into the compiler.
// Does not compile.

package main

import "runtime"

func main() {
	runtime.printbool(true)	// ERROR "unexported|undefined"
}
