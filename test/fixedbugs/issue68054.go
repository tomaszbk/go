// compile

package p

type Seq[V any] = func(yield func(V) bool)

func f[E any](seq Seq[E]) {
	return
}

func g() {
	f(Seq[int](nil))
}

type T[P any] struct{}

type A[P any] = T[P]

var _ A[int]
