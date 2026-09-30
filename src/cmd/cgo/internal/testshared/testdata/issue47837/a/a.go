package a

type A interface {
	M()
}

//go:noinline
func TheFuncWithArgA(a A) {
	a.M()
}

type ImplA struct{}

//go:noinline
func (A *ImplA) M() {}
