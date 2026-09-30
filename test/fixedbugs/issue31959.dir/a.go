package a

type T struct{}

func F() {
	type T = int
	println(T(0))
}
