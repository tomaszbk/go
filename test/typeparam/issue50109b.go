// run

package main

func main() {
	F[any]()
}

func F[T any]() I[T] {
	return (*S1[T])(nil)
}

type I[T any] interface{}

type S1[T any] struct {
	*S2[T]
}

type S2[T any] struct {
	S3 *S3[T]
}

type S3[T any] struct {
	x int
}
