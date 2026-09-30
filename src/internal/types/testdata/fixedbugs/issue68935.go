package p

type A = struct {
	F string
	G int
}

func Make[T ~A]() T {
	return T{
		F: "blah",
		G: 1234,
	}
}

type N struct {
	F string
	G int
}

func _() {
	_ = Make[N]()
}
