// run

package main

type myifacer[T any] interface{ do(T) error }

type stuff[T any] struct{}

func (s stuff[T]) run() interface{} {
	var i myifacer[T]
	return i
}

func main() {
	stuff[int]{}.run()
}
