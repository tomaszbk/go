package p

var m map[int]int

func _() {
	_, ok /* ERROR "undefined: ok" */ = m[0] // must not crash
}

func _() {
	var ok = undef /* ERROR "undefined: undef" */
	x, ok := m[0]  // must not crash
	_, _ = x, ok
}
