// run


// Test that interface conversion fails when method is missing.

package main

type I interface {
	Foo()
}

func main() {
	shouldPanic(p1)
}

func p1() {
	var s *S
	var i I
	var e interface{}
	e = s
	i = e.(I)
	_ = i
}

type S struct{}

func (s *S) _() {}

func shouldPanic(f func()) {
	defer func() {
		if recover() == nil {
			panic("function should panic")
		}
	}()
	f()
}
