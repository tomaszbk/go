package b

type W struct{}

var Called bool

func (W) M() {
	Called = true
}
