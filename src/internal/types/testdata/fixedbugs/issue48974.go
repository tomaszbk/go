package p

type Fooer interface {
	Foo()
}

type Fooable[F /* ERROR "instantiation cycle" */ Fooer] struct {
	ptr F
}

func (f *Fooable[F]) Adapter() *Fooable[*FooerImpl[F]] {
	return &Fooable[*FooerImpl[F]]{&FooerImpl[F]{}}
}

type FooerImpl[F Fooer] struct {
}

func (fi *FooerImpl[F]) Foo() {}
