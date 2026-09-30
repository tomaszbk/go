// errorcheck -goexperiment fieldtrack


package main

type Fooer interface {
	Foo() string
}

type FooImpl struct{}

//go:nointerface
func (FooImpl) Foo() string { return "foo" }

func toInterface[T Fooer](fooer T) Fooer {
	return fooer
}

func main() {
	var iface Fooer = toInterface(FooImpl{}) // ERROR "does not satisfy Fooer"
	iface.Foo()
}
