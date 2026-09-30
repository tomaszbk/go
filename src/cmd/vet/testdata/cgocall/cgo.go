// This file contains tests for the cgo checker.

package testdata

// void f(void *p) {}
import "C"

import "unsafe"

func CgoTests() {
	var c chan bool
	C.f(*(*unsafe.Pointer)(unsafe.Pointer(&c))) // ERROR "embedded pointer"
	C.f(unsafe.Pointer(&c))                     // ERROR "embedded pointer"
}
