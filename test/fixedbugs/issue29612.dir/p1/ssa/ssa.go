package ssa

type T struct{}

func (T) foo() {}

type fooer interface {
	foo()
}

func Unused(v interface{}) {
	v.(fooer).foo()
	v.(interface{ foo() }).foo()
}
