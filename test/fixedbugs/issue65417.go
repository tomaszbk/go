// run

package main

import (
	"strings"
	"unsafe"
)

func main() {
	shouldPanic("runtime error: index out of range", func() { f(0) })
	shouldPanic("runtime error: index out of range", func() { g(0) })
}

func f[T byte](t T) {
	const str = "a"
	_ = str[unsafe.Sizeof(t)]
}

func g[T byte](t T) {
	const str = "a"
	_ = str[unsafe.Sizeof(t)+0]
}

func shouldPanic(str string, f func()) {
	defer func() {
		err := recover()
		if err == nil {
			panic("did not panic")
		}
		s := err.(error).Error()
		if !strings.Contains(s, str) {
			panic("got panic " + s + ", want " + str)
		}
	}()

	f()
}
