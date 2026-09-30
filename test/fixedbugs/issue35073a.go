// run -gcflags=-d=checkptr


// Test that reflect.Value.UnsafeAddr/Pointer is handled
// correctly by -d=checkptr

package main

import (
	"reflect"
	"unsafe"
)

func main() {
	n := 10
	m := make(map[string]string)

	_ = unsafe.Pointer(reflect.ValueOf(&n).Elem().UnsafeAddr())
	_ = unsafe.Pointer(reflect.ValueOf(&m).Elem().Pointer())
}
