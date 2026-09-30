package target

type Base struct{}

func (Base) M() {}

type Target struct {
	Base
}

var P *Target
