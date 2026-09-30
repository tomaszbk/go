package p

import "unsafe"

type T /* ERROR "invalid recursive type" */ [unsafe.Sizeof(f())]int

func f() T {
    return T{}
}