// errorcheck


// Test that a syntax error caused by an unexpected EOF
// gives an error message with the correct line number.
//
// https://golang.org/issue/3392

package main

func foo() {
	bar(1, // ERROR "unexpected|missing|undefined"