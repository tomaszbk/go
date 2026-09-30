// errorcheck -0 -m -l


package foo

// Escape analysis needs to treat the uintptr-typed reflect.*Header fields as pointers.

import (
	"reflect"
	"unsafe"
)

type immutableBytes []byte

// Bug was failure to leak param b.
func toString(b immutableBytes) string { // ERROR "leaking param: b$"
	var s string
	if len(b) == 0 {
		return s
	}

	strHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	strHeader.Data = (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data

	l := len(b)
	strHeader.Len = l
	return s
}
