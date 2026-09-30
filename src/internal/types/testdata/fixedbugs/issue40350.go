package p

type number interface {
	~float64 | ~int | ~int32
	float64 | ~int32
}

func f[T number]() {}

func _() {
	_ = f[int /* ERROR "int does not satisfy number (number mentions int, but int is not in the type set of number)" */]
}
