// run

package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	defer func() {
		e := recover()
		if e == nil {
			panic("should have panicked")
		}
		text := fmt.Sprintf("%s", e) // handles both string and runtime.errorString
		if !strings.HasPrefix(text, "reflect:") {
			panic("wanted a reflect error, got this instead:\n" + text)
		}
	}()
	r := reflect.MakeFunc(reflect.TypeOf(func() error { return nil }),
		func(args []reflect.Value) []reflect.Value {
			var x [1]reflect.Value
			return x[:]
		}).Interface().(func() error)
	r()
}
