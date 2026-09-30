// Test that a linkname applied on an assembly declaration
// does not affect stack map generation.

package main

import (
	"runtime"
	_ "unsafe"
)

//go:linkname asm
func asm(*int)

func main() {
	x := new(int)
	asm(x)
}

// called from asm
func callback() {
	runtime.GC() // scan stack
}
