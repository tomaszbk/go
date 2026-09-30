package p

type I[T any] interface {
	m(I[T])
}

type S[T any] struct{}

func (S[T]) m(I[T]) {}

func f[T I[E], E any](T) {}

func _() {
	f(S[int]{})
}
