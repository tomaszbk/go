package a

import "unsafe"

func ConstUnsafePointer() unsafe.Pointer {
	return unsafe.Pointer(uintptr(0))
}
