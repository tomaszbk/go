package a

func F[T any]() A[T] {
	var x A[T]
	return x
}

type A[T any] struct {
	b B[T]
}

func (a A[T]) M() C[T] {
	return C[T]{
		B: a.b,
	}
}

type B[T any] struct{}

type C[T any] struct {
	B B[T]
}
