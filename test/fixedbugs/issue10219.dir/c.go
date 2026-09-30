package c

import "./b"

func F() {
	s := b.F()
	s.M("c")
}
