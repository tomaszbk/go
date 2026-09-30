package b

import "./a"

func F() interface{} {
	return a.F()
}

func P() interface{} {
	return a.P()
}
