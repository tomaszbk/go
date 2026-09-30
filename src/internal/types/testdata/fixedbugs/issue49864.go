package p

func _[P ~int, Q any](p P) {
	_ = Q(p /* ERROR "cannot convert" */ )
}
