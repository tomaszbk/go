// errorcheck


package main

import "unsafe"

const (
	_ = unsafe.Sizeof([0]byte{}[0]) // ERROR "out of bounds"
)
