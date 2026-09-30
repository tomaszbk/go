// run

// Similar to reflectmethod5.go, but for reflect.Type.MethodByName.

package main

import "reflect"

var called bool

type foo struct{}

func (foo) X() { called = true }

var h = reflect.Type.MethodByName

func main() {
	v := reflect.ValueOf(foo{})
	m, ok := h(v.Type(), "X")
	if !ok {
		panic("FAIL")
	}
	f := m.Func.Interface().(func(foo))
	f(foo{})
	if !called {
		panic("FAIL")
	}
}
