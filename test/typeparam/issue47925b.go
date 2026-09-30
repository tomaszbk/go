// run


package main

type I[T any] interface {
	foo()
}

type E[T any] interface {
}

//go:noinline
func f[T I[T]](x T) E[T] {
	// contains a cast from nonempty to empty interface
	return E[T](I[T](x))
}

type S struct {
	x int
}

func (s *S) foo() {}

func main() {
	i := f(&S{x: 7})
	if i.(*S).x != 7 {
		panic("bad")
	}
}
