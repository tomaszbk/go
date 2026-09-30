package p

import _ "unsafe"

// f1 is pushed from main.
//
//go:linkname f1
func f1()

// Push f2 to main.
//
//go:linkname f2 main.f2
func f2() {}

func F() { f1() }
