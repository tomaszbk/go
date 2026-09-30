package p

func _[P int | float64 | complex128]() {
	_ = map[P]int{1: 1, 1.0 /* ERROR "duplicate key 1" */ : 2, 1 /* ERROR "duplicate key (1 + 0i)" */ + 0i: 3}
}
