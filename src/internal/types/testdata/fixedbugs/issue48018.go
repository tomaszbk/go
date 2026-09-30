package main

type Box[A any] struct {
	value A
}

func Nest[A /* ERROR "instantiation cycle" */ any](b Box[A], n int) interface{} {
	if n == 0 {
		return b
	}
	return Nest(Box[Box[A]]{b}, n-1)
}

func main() {
	Nest(Box[int]{0}, 10)
}
