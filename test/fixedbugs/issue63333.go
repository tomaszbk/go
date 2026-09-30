// errorcheck -goexperiment fieldtrack


package p

func f(interface{ m() }) {}
func g()                 { f(new(T)) } // ERROR "m method is marked 'nointerface'"

type T struct{}

//go:nointerface
func (*T) m() {}
