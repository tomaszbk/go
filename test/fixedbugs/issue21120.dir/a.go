package a

type S struct {
	x int
}

func V() interface{} {
	return S{0}
}
