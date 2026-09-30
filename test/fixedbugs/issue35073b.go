// errorcheck -0 -d=checkptr -m

// Test that we can inline the receiver arguments for
// reflect.Value.UnsafeAddr/Pointer, even in checkptr mode.

package main

import (
	"reflect"
	"unsafe"
)

func main() {
	n := 10                      // ERROR "moved to heap: n"
	m := make(map[string]string) // ERROR "moved to heap: m" "make\(map\[string\]string\) escapes to heap"

	_ = unsafe.Pointer(reflect.ValueOf(&n).Elem().UnsafeAddr()) // ERROR "inlining call"
	_ = unsafe.Pointer(reflect.ValueOf(&m).Elem().Pointer())    // ERROR "inlining call"
}
