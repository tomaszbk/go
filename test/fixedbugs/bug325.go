// errorcheck


package main

import "unsafe"

func main() {
	var x unsafe.Pointer
	println(*x) // ERROR "invalid indirect.*unsafe.Pointer|cannot indirect"
	var _ = (unsafe.Pointer)(nil).foo  // ERROR "foo"
}
