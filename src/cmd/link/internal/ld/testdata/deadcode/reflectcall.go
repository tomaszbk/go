// This example uses reflect.Value.Call, but not
// reflect.{Value,Type}.Method. This should not
// need to bring all methods live.

package main

import "reflect"

func f() { println("call") }

type T int

func (T) M() {}

func main() {
	v := reflect.ValueOf(f)
	v.Call(nil)
	i := interface{}(T(1))
	println(i)
}
