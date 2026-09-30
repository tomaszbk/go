package b

import "./a"

func f() {
	a.Marshal(map[int]int{})
}
