package b

import "./a"

//go:noinline
func F() interface{} {
	return a.T[int]{}
}

//go:noinline
func G() interface{} {
	return struct{ X, Y a.U }{}
}
