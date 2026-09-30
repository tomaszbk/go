// errorcheck

// Tests that errors are reported for misuse of linkname.
package p

import _ "unsafe"

type t int

var x, y int

func F[T any](T) {}

//go:linkname x ok

// ERROR "//go:linkname must refer to declared function or variable"
// ERROR "//go:linkname must refer to declared function or variable"
// ERROR "duplicate //go:linkname for x"
// ERROR "//go:linkname reference of an instantiation is not allowed"

//line linkname3.go:16
//go:linkname nonexist nonexist
//go:linkname t notvarfunc
//go:linkname x duplicate
//go:linkname i F[go.shape.int]
