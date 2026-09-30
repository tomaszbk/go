// compile

package p

import "unsafe"

func F[T int](v T) uintptr {
	return unsafe.Offsetof(struct{ f T }{
		func(T) T { return v }(v),
	}.f)
}

func f() {
	F(1)
}
