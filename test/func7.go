// run

// Test evaluation order in if condition.

package main

var calledf = false

func f() int {
	calledf = true
	return 1
}

func g() int {
	if !calledf {
		panic("BUG: func7 - called g before f")
	}
	return 0
}

func main() {
	// gc used to evaluate g() before f().
	if f() < g() {
		panic("wrong answer")
	}
}
