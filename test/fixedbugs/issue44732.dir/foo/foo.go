package foo

type Foo struct {
	updatecb func()
}

func NewFoo() *Foo {
	return &Foo{updatecb: nil}
}
