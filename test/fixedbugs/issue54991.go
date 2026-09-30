// compile


package p

import (
	"sync/atomic"
)

type I interface {
	M()
}

type S struct{}

func (*S) M() {}

type T struct {
	I
	x atomic.Int64
}

func F() {
	t := &T{I: &S{}}
	t.M()
}
