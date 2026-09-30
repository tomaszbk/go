// run -gcflags=-d=checkptr


package main

import (
	"reflect"
	"unsafe"
)

var s []int

func main() {
	s = []int{42}
	h := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	x := *(*int)(unsafe.Pointer(h.Data))
	if x != 42 {
		panic(x)
	}
}
