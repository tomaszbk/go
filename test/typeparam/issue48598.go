// run

package main

type Iterator[T any] interface {
	Iterate()
}

type IteratorFunc[T any] func(fn func(T) bool)

func (f IteratorFunc[T]) Iterate() {
}

func FromIterator[T any](it Iterator[T]) {
	it.Iterate()
}

func Foo[T, R any]() {
	FromIterator[R](IteratorFunc[R](nil))
}

func main() {
	Foo[int, int]()
}
