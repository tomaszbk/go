// Linkname systemstack is not allowed, even if it is
// defined in assembly.

package main

import _ "unsafe"

func f() {}

func main() {
	systemstack(f)
}

//go:linkname systemstack runtime.systemstack
func systemstack(func())
