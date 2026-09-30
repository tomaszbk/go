// compile


package main

import "reflect"

type E struct{}

func (E) M() {}

func g(x struct{ a, b int }) {
	reflect.ValueOf(x)
}

func gEmbedded(x struct {
	E
	b int
}) {
	reflect.ValueOf(x)
}

func f[T any](i interface{}) {
	switch i.(type) {
	case T:
	case struct{ a, b T }:
	}
}

func fEmbedded[T any](i interface{}) {
	switch i.(type) {
	case T:
	case struct {
		E
		b T
	}:
	}
}

func main() {
	f[int](0)
	f[any](0)
	fEmbedded[int](0)
	fEmbedded[any](0)
}
