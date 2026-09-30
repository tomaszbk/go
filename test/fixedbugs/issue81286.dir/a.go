package a

type W struct{}

var Last int

func (W) M(x int) {
	Last = x
}
