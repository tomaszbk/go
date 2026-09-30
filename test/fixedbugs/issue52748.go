// errorcheck


package p

import "unsafe"

type S[T any] struct{}

const c = unsafe.Sizeof(S[[c]byte]{}) // ERROR "initialization cycle"
