// compile


package main

type F func() F

func do[T any]() F {
	return nil
}

type G[T any] func() G[T]

//go:noinline
func dog[T any]() G[T] {
	return nil
}

func main() {
	do[int]()
	dog[int]()
}
