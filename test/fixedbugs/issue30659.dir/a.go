package a

type I interface {
	I2
}
type I2 interface {
	M()
}
type S struct{}

func (*S) M() {}

func New() I {
	return &S{}
}
