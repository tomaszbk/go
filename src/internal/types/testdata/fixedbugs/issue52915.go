package p

import "unsafe"

type T[P any] struct {
	T /* ERROR "invalid recursive type" */ [P]
}

func _[P any]() {
	_ = unsafe.Sizeof(T[int]{})
	_ = unsafe.Sizeof(struct{ T[int] }{})

	_ = unsafe.Sizeof(T[P]{})
	_ = unsafe.Sizeof(struct{ T[P] }{})
}

const _ = unsafe /* ERROR "not constant" */ .Sizeof(T /* ERROR "invalid recursive type" */ [int]{})
