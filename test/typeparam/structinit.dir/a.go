package a

type S[T any] struct {
}

func (b *S[T]) build() *X[T] {
	return &X[T]{f:0}
}
type X[T any] struct {
	f int
}
