// run -race

//go:build cgo && linux && amd64

package main

import (
	"fmt"
	"testing"
	"unsafe"
)

var buf [2]byte
var x unsafe.Pointer = unsafe.Pointer(&buf[0])

func main() {
	n := testing.AllocsPerRun(1000, func() {
		x = unsafe.Pointer(uintptr(x) + 1)
		x = unsafe.Pointer(uintptr(x) - 1)
	})
	if n > 0 {
		panic(fmt.Sprintf("too many allocations; want 0 got %f", n))
	}
}
