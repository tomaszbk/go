package p

func _() {
	var x int
	var f func() []int
	_ = f /* ERROR "cannot index f" */ [x]
}
