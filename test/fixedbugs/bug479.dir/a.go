package p

import "unsafe"

type S2 struct {}

const C = unsafe.Sizeof(S2{})

type S1 struct {
	S2
}
