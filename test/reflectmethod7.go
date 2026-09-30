// run


// See issue 44207.

package main

import "reflect"

type S int

func (s S) M() {}

func main() {
	t := reflect.TypeOf(S(0))
	fn, ok := reflect.PointerTo(t).MethodByName("M")
	if !ok {
		panic("FAIL")
	}
	fn.Func.Call([]reflect.Value{reflect.New(t)})
}
