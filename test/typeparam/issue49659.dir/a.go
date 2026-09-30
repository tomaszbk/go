package a

type A[T any] struct {
	a int
}

func (a A[T]) F() {
	_ = &a.a
}
