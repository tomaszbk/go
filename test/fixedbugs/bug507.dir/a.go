package a

type I interface {
	M()
}

type S struct {
	I I
}
