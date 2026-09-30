package p

type S struct{}

func (S) M() byte {
	return 0
}

type I[T any] interface {
	M() T
}

func f[T any](x I[T]) {}

func _() {
	f(S{})
}
