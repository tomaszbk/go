// compile


package p

type F = func(T)

type T interface {
	m(F)
}

type t struct{}

func (t) m(F) {}

var _ T = &t{}
