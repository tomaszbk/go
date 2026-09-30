// run

package main

type Constraint[T any] interface {
	~func() T
}

func Foo[T Constraint[T]]() T {
	var t T

	t = func() T {
		return t
	}
	return t
}

func main() {
	type Bar func() Bar
	Foo[Bar]()
}
