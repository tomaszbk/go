package b

import "./a"

func B() string {
	return a.M()
}
