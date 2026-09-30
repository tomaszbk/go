// errorcheck


// Issue 2627 -- unsafe.Pointer type isn't handled nicely in some errors

package main

import "unsafe"

func main() {
	var x *int
	_ = unsafe.Pointer(x) - unsafe.Pointer(x) // ERROR "(operator|operation) - not defined on unsafe.Pointer|expected integer, floating, or complex type"
}
