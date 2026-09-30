package p2

import "./p1"

type T struct {
	p1.T
}

func F() {
	var t T
	p1.F(&t.T)
}
