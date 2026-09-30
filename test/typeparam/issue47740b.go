// run

package main

import "reflect"

type S[T any] struct {
	a interface{}
}

func (e S[T]) M() {
	v := reflect.ValueOf(e.a)
	_, _ = v.Interface().(int)
}

func main() {
	e := S[int]{0}
	e.M()
}
