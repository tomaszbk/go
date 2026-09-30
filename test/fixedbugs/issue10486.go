// run


// Issue 10486.
// Check stack walk during div by zero fault,
// especially on software divide systems.

package main

import "runtime"

var A, B int

func divZero() int {
	defer func() {
		if p := recover(); p != nil {
			var pcs [512]uintptr
			runtime.Callers(2, pcs[:])
			runtime.GC()
		}
	}()
	return A / B
}

func main() {
	A = 1
	divZero()
}
