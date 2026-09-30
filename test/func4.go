// errorcheck


// Verify that it is illegal to take the address of a function.
// Does not compile.

package main

var notmain func()

func main() {
	var x = &main		// ERROR "address of|invalid"
	main = notmain	// ERROR "assign to|invalid"
	_ = x
}
