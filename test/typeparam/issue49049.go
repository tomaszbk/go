// run


package main

type A[T any] interface {
	m()
}

type Z struct {
	a,b int
}

func (z *Z) m() {
}

func test[T any]() {
	var a A[T] = &Z{}
	f := a.m
	f()
}
func main() {
	test[string]()
}
