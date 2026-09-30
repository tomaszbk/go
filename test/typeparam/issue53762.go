// compile

package main

type Value[T any] interface {
}

func use[T any](v Value[T]) {
	_, _ = v.(int)
}

func main() {
	use[int](Value[int](1))
}
