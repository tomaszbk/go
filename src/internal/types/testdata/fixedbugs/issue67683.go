package p

type A[P any] func()

// alias signature types
type B[P any] = func()
type C[P any] = B[P]

var _ = A /* ERROR "cannot use generic type A without instantiation" */ (nil)

// generic alias signature types must be instantiated before use
var _ = B /* ERROR "cannot use generic type B without instantiation" */ (nil)
var _ = C /* ERROR "cannot use generic type C without instantiation" */ (nil)
