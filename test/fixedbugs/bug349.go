// errorcheck


// issue 1192 - detail in error

package main

func foo() (a, b, c int) {
	return 0, 1 2.01  // ERROR "unexpected literal 2.01|expected ';' or '}' or newline|not enough arguments to return"
}
