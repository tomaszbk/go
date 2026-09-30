package p

func f[P, Q any]() {}

// Don't panic on the next line, produce a proper error instead.
var _ any = f /* ERROR "cannot infer Q" */ [int]
