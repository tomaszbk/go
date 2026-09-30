// run fake-arg-to-force-use-of-go-run


//go:build cgo

package main

// void f(int *p) { *p = 0x12345678; }
import "C"

func main() {
	var x C.int
	func() {
		defer C.f(&x)
	}()
	if x != 0x12345678 {
		panic("FAIL")
	}
}
