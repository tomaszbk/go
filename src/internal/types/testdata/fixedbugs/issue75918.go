package p

import "unsafe"

type A /* ERROR "invalid recursive type" */ [unsafe/* ERROR "must be constant" */.Sizeof(S{})]byte

type S struct {
	a A
}
