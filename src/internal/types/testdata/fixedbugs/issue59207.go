package p

import "unsafe"

type E [1 << 32]byte

var a [1 << 32]E // size of a must not overflow to 0
var _ = unsafe.Sizeof(a /* ERROR "too large" */ )
