package p

func f() {}

const c = 0

var v int
var _ = f < c // ERROR "invalid operation: f < c (mismatched types func() and untyped int)"
var _ = f < v // ERROR "invalid operation: f < v (mismatched types func() and int)"
