package p

import "unsafe"

type T /* ERROR "invalid recursive type: T refers to itself" */ struct {
	T
}

func _(t T) {
	_ = unsafe.Sizeof(t) // should not go into infinite recursion here
}
