package a

type A[T any] struct {
}

func (a A[T]) F() {
	_ = a
}
