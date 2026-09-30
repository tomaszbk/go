package c

import "./a"

func C() {
	var x int64
	println(a.F(&x, &x))
	var y int32
	println(a.F(&y, &y))
}
