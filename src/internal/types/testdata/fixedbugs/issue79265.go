package p

import "unsafe"

type _ A // force type-checking of A first; order is relevant
type A /* ERROR "invalid recursive type A" */ = T
type T [unsafe.Sizeof(A{})]int

// same as above, but through a chain of aliases
type _ A1
type A1 /* ERROR "invalid recursive type A1" */ = A2
type A2 = U
type U [unsafe.Sizeof(A1{})]int
