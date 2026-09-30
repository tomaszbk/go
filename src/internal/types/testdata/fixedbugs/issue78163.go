package p

func _[P any](x P) {
	var s []byte
	_ = append(s, x /* ERROR "cannot use x (variable of type P constrained by any) as []byte value in argument to append" */ ...)
	copy(s, x /* ERROR "invalid copy: argument must be a slice; have x (variable of type P constrained by any)" */)
}
