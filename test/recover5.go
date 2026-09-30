// errorcheck

// Verify that recover arguments requirements are enforced by the
// compiler.

package main

func main() {
	_ = recover()     // OK
	_ = recover(1)    // ERROR "too many arguments"
	_ = recover(1, 2) // ERROR "too many arguments"
}
