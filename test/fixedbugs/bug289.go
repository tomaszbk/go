// errorcheck


// https://code.google.com/p/gofrontend/issues/detail?id=1

package main

func f1() {
	a, b := f() // ERROR "assignment mismatch|does not match|cannot initialize"
	_, _ = a, b
}

func f2() {
	var a, b int
	a, b = f() // ERROR "assignment mismatch|does not match|cannot assign"
	_, _ = a, b
}

func f() int {
	return 1
}
