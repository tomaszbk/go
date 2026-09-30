// errorcheck

package main

func foo() (T, T) { // ERROR "undefined"
	return 0, 0
}

func bar() (T, string, T) { // ERROR "undefined"
	return 0, "", 0
}

func main() {
	var x, y, z int
	x, y = foo()
	x, y, z = bar() // ERROR "cannot (use type|assign|use.*type) string|"
	_, _, _ = x, y, z
}
