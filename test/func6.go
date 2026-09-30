// run


// Test closures in if conditions.

package main

func main() {
	if func() bool { return true }() {}  // gc used to say this was a syntax error
	if (func() bool { return true })() {}
	if (func() bool { return true }()) {}
}

