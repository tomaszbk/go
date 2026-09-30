// compile

package main

type B[T any] struct {
	a A[T]
}

type A[T any] = func(B[T]) bool

func main() {
	var s A[int]
	println(s)
}
