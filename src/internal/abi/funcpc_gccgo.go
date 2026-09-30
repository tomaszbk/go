// For bootstrapping with gccgo.

//go:build gccgo

package abi

import "unsafe"

func FuncPCABI0(f interface{}) uintptr {
	words := (*[2]unsafe.Pointer)(unsafe.Pointer(&f))
	return *(*uintptr)(unsafe.Pointer(words[1]))
}

func FuncPCABIInternal(f interface{}) uintptr {
	words := (*[2]unsafe.Pointer)(unsafe.Pointer(&f))
	return *(*uintptr)(unsafe.Pointer(words[1]))
}
