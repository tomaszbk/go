package b

import "./a"

func f() {
	for k := range (a.M{}) {
		k.F()
	}
}
