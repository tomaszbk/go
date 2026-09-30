package p

func _() {
	_ = m[] // ERROR expected operand
	_ = m[x,]
	_ = m[x /* ERROR unexpected name a */ a b c d]
}

// test case from the issue
func f(m map[int]int) int {
	return m[0 // ERROR expected comma, \: or \]
		]
}
