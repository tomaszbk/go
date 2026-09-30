package a

type dO struct {
	x int
}

type EDO struct{}

func (EDO) Apply(*dO) {}

var X EDO
