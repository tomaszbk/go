// compile

package main

type Interface[T any] interface {
}

func F[T any]() Interface[T] {
	var i int
	return i
}

func main() {
	F[int]()
}
