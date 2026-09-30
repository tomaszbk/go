package p

type I1 interface {
	Foo(int)
}

type I2 interface {
	foo(int)
}

type M1 int

func (M1) foo() {}

type M2 int

func (M2) foo(int) {}
