package p3

import "./p2"

func F() {
	p2.F()
	var t p2.T
	println(t.T.M())
}
