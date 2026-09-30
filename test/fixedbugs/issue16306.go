// compile

package main

import "unsafe"

var x = unsafe.Pointer(uintptr(0))

func main() {
	_ = map[unsafe.Pointer]int{unsafe.Pointer(uintptr(0)): 0}
}
