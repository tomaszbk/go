package ssa

type T struct{}

func (T) foo() {}

type fooer interface {
	foo()
}

func Works(v interface{}) {
	switch v.(type) {
	case interface{}:
		v.(fooer).foo()
	}
}

func Panics(v interface{}) {
	switch v.(type) {
	case interface{}:
		v.(fooer).foo()
		v.(interface{ foo() }).foo()
	}
}
