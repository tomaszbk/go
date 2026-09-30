package a0

type Builder[T any] struct{}

func (r Builder[T]) New1() T {
	var v T
	return v
}

func (r Builder[T]) New2() T {
	var v T
	return v
}

type IntBuilder struct{}

func (b IntBuilder) New() int {
	return Builder[int]{}.New2()
}
