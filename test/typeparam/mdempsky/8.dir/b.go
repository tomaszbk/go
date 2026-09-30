package b

import "./a"

func init() {
	a.F[func()]() // ERROR "does not satisfy comparable"
}
