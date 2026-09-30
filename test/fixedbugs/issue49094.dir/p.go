package p

import (
	"./b"
)

type S struct{}

func (S) M() {
	b.M(nil)
}
