package bug

type S struct {
	F func()
}

type X interface {
	Bar()
}

func Foo(x X) *S {
	return &S{F: x.Bar}
}
