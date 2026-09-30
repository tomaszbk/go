// run


package main

type A[T any] struct {
	field B[T]
}

type B[T any] interface {
	Work(T)
}

func (a *A[T]) Work(t T) {
	a.field.Work(t)
}

type BImpl struct{}

func (b BImpl) Work(s string) {}

func main() {
	a := &A[string]{
		field: BImpl{},
	}
	a.Work("")
}
