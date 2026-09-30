// build


package main

type Foo[T any] interface {
	CreateBar() Bar[T]
}

type Bar[T any] func() Bar[T]

func (f Bar[T]) CreateBar() Bar[T] {
	return f
}

func abc[R any]() {
	var _ Foo[R] = Bar[R](nil)()
}

func main() {
	abc[int]()
}