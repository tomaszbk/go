// run

package main

import (
	"reflect"
	"runtime"
)

type T struct {
	a, b int
}

func f(t *T) int {
	if t != nil {
		return t.b
	}
	return 0
}

func g(t *T) int {
	return f(t) + 5
}

func main() {
	x(f)
	x(g)
}
func x(v any) {
	println(runtime.FuncForPC(reflect.ValueOf(v).Pointer()).Name())
}
