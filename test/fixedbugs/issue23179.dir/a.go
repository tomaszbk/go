package a

type Large struct {
	x [256]int
}

func F(x int, _ int, _ bool, _ Large) int {
	return x
}
