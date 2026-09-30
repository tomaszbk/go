// run

package main

import (
	"reflect"
	"runtime"
)

func f(n int) int {
	return n % 2
}

func g(n int) int {
	return f(n)
}

func name(fn any) (res string) {
	return runtime.FuncForPC(uintptr(reflect.ValueOf(fn).Pointer())).Name()
}

func main() {
	println(name(f))
	println(name(g))
}
