// run

//go:build cgo

package main

import (
	"log"
	"runtime/cgo"
	"unsafe"
)

type S struct{ _ cgo.Incomplete }

func main() {
	p := (*S)(unsafe.Pointer(uintptr(0x8000)))
	var v any = p
	p2 := v.(*S)
	if p != p2 {
		log.Fatalf("%p != %p", unsafe.Pointer(p), unsafe.Pointer(p2))
	}
	p2 = typeAssert[*S](v)
	if p != p2 {
		log.Fatalf("%p != %p from typeAssert", unsafe.Pointer(p), unsafe.Pointer(p2))
	}
}

func typeAssert[T any](v any) T {
	return v.(T)
}
