package p

import "unsafe"

type A /* ERROR "invalid recursive type" */ [unsafe.Sizeof(B{})]int
type B A

type C /* ERROR "invalid recursive type" */ [unsafe.Sizeof(f())]int
func f() D {
	return D{}
}
type D C

type E /* ERROR "invalid recursive type" */ [unsafe/* ERROR "must be constant" */.Sizeof(g[F]())]int
func g[P any]() P {
	panic(0)
}
type F struct {
	f E
}

