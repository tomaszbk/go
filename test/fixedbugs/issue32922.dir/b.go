package b

import "./a"

func B() int {
	return 99 + a.A()
}
