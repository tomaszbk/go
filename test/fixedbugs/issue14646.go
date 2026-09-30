// run


package main

import "runtime"

func main() {
	var file string
	var line int
	func() {
		defer func() {
			_, file, line, _ = runtime.Caller(1)
		}()
	}() // this is the expected line
	const EXPECTED = 15
	if line != EXPECTED {
		println("Expected line =", EXPECTED, "but got line =", line, "and file =", file)
	}
}
