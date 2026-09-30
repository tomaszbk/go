// compile


package p

import "unsafe"

func F[T any](v T) uintptr {
	return unsafe.Alignof(func() T {
		func(any) {}(struct{ _ T }{})
		return v
	}())
}

func f() {
	F(0)
}
