// run

package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var s = "abc"
	sh1 := (*reflect.StringHeader)(unsafe.Pointer(&s))
	ptr2 := unsafe.Pointer(unsafe.StringData(s))
	if ptr2 != unsafe.Pointer(sh1.Data) {
		panic(fmt.Errorf("unsafe.StringData ret %p != %p", ptr2, unsafe.Pointer(sh1.Data)))
	}
}
