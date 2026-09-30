package b

import "./a"

func G(x int) int {
	return a.F(x, 1, false, a.Large{})
}
