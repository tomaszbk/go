package b

import "./a"

type S struct {
	a.I
}

var V a.I

func init() {
	V = S{}
}
