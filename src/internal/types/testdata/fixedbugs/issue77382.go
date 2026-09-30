package p

import "unsafe"

type T1[T any] [42]T2[T]

type T2[T any] [42]T

func _[T any]() {
	_ = unsafe.Sizeof(T1[T]{})
}
