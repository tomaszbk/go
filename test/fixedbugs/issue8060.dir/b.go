package b

import "./a"

var X = a.A

func b() {
	_ = [3][1]float64{}
}
