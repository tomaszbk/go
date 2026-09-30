// compile

// Issue 1787.

package main

import "unsafe"

const x = unsafe.Sizeof([8]byte{})

func main() {
	var b [x]int
	_ = b
}

/*
bug338.go:14: array bound must be non-negative
*/
