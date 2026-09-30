package p

func A(x interface {
	X() int
}) int {
	return x.X()
}

func B(x interface {
	X() int
}) int {
	return x.X()
}
