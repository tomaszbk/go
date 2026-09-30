package main

import (
	_ "unsafe"

	"./a"
)

//go:linkname s test/a.s
var s string

func main() {
	if a.Get() != "a" {
		panic("FAIL")
	}

	s = "b"
	if a.Get() != "b" {
		panic("FAIL")
	}
}
