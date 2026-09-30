package main

import "./lib"

type I interface {
	M()
}

type PI interface {
	PM()
}

func main() {
	var t lib.T
	t.M()
	t.PM()

	// This is still an error.
	// var i1 I = t
	// i1.M()
	
	// This combination is illegal because
	// PM requires a pointer receiver.
	// var pi1 PI = t
	// pi1.PM()

	var pt = &t
	pt.M()
	pt.PM()

	var i2 I = pt
	i2.M()

	var pi2 PI = pt
	pi2.PM()
}
