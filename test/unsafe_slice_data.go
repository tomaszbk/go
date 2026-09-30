// run


package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var s = []byte("abc")
	sh1 := *(*reflect.SliceHeader)(unsafe.Pointer(&s))
	ptr2 := unsafe.Pointer(unsafe.SliceData(s))
	if ptr2 != unsafe.Pointer(sh1.Data) {
		panic(fmt.Errorf("unsafe.SliceData %p != %p", ptr2, unsafe.Pointer(sh1.Data)))
	}
}
