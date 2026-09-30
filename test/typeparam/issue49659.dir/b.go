package b

import "./a"

type B[T any] struct {
	v a.A[T]
}

func (b B[T]) F() {
	b.v.F()
}
