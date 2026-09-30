package p

import "unsafe"

func _[T any](x T) {
	_ = unsafe.Alignof(x)
	_ = unsafe.Sizeof(x)
}
