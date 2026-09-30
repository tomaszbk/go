package a

type A[T any] struct{}

func (_ A[T]) Method() {}

func DoSomething[P any]() {
	a := A[*byte]{}
	a.Method()
}
