// compile


package p

type B[T any] interface {
	Work()
}
type BImpl[T any] struct{}

func (b *BImpl[T]) Work() {
}

type A[T any] struct {
	B[T]
}

func f[T any]() {
	s := &A[T]{
		&BImpl[T]{},
	}
	// golang.org/issue/48056
	s.Work()
}
