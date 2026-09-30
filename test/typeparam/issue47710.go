// compile


package p

type FooType[t any] interface {
	Foo(BarType[t])
}
type BarType[t any] interface {
	Int(IntType[t]) FooType[int]
}

type IntType[t any] int

func (n IntType[t]) Foo(BarType[t]) {}
func (n IntType[_]) String()    {}
