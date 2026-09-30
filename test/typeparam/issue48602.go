// run


package main

type Iterator[T any] interface {
	Iterate(fn T)
}

type IteratorFunc[T any] func(fn T)

func (f IteratorFunc[T]) Iterate(fn T) {
	f(fn)
}

func Foo[R any]() {
	var _ Iterator[R] = IteratorFunc[R](nil)
}

func main() {
	Foo[int]()
}
