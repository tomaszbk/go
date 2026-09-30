package p

import "unsafe"

func F1[T any](_ [unsafe.Sizeof(F1[int])]T) (res T)      { return }
func F2[T any](_ T) (res [unsafe.Sizeof(F2[string])]int) { return }
func F3[T any](_ [unsafe.Sizeof(F1[string])]int)         {}
