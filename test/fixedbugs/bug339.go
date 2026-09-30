// run

// Issue 1608.
// Size used to be -1000000000.

package main

import "unsafe"

func main() {
	var a interface{} = 0
	size := unsafe.Sizeof(a)
	if size != 2*unsafe.Sizeof((*int)(nil)) {
		println("wrong size: ", size)
	}
}
