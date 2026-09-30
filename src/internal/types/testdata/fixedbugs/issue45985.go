package issue45985

func app[S interface{ ~[]T }, T any](s S, e T) S {
	return append(s, e)
}

func _() {
	_ = app /* ERROR "S (type int) does not satisfy interface{~[]T}" */ [int] // TODO(gri) better error message
}
