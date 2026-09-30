// errorcheck


// Verify that erroneous use of init is detected.
// Does not compile.

package main

func init() {
}

func main() {
	init()         // ERROR "undefined.*init"
	runtime.init() // ERROR "undefined.*runtime\.init|reference to undefined name|undefined: runtime"
	var _ = init   // ERROR "undefined.*init"
}
