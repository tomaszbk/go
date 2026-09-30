// run


package main

import (
	"reflect"
	"unsafe"
)

func main() {
	t := reflect.TypeOf(unsafe.Pointer(nil))
	if pkgPath := t.PkgPath(); pkgPath != "unsafe" {
		panic("unexpected t.PkgPath(): " + pkgPath)
	}
}
