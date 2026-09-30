// compile


package p

type I interface {
	M(interface{})
}

type a[T any] struct{}

func (a[T]) M(interface{}) {}

func f[T I](t *T) {
	(*t).M(t)
}

func g() {
	f(&a[int]{})
}
