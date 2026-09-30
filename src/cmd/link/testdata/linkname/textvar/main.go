// Using a linknamed variable to reference an assembly
// function in the same package is ok.

package main

import _ "unsafe"

func main() {
	println(&asmfunc)
}

//go:linkname asmfunc
var asmfunc uintptr
