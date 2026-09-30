package p

func _() {
	const x = 0
	x /* ERROR "cannot assign to x" */ += 1
	x /* ERROR "cannot assign to x" */ ++
}
