package p

type A[T any] int

func (A[T]) m(A[T])

func f[P interface{m(P)}]() {}

func _() {
	_ = f[A[int]]
}
