package q

import "./p"

type T struct{}

func (T) M() interface{} {
	return &p.T{}
}
