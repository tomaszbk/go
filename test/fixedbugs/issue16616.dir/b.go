package b

import "./a"

var V struct{ i int }

var U struct {
	a.V
	j int
}
