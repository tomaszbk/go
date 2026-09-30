package a

type T2 struct{}

func (t *T2) M2(a, b float64) {
	variadic(a, b)
}

func variadic(points ...float64) {
	println(points)
}
