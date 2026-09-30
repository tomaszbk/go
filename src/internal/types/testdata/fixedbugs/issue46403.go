package issue46403

func _() {
	// a should be used, despite the parser error below.
	var a []int
	var _ = a[] // ERROR "expected operand"
}
