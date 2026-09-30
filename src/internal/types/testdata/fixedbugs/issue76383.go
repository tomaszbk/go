package p

import "unsafe"

var v any = 42

type T /* ERROR "invalid recursive type" */ struct {
	f [unsafe.Sizeof(v.(T))]int
}
