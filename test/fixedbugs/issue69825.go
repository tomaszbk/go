// compile -d=libfuzzer

package main

type T struct {
	A
}

type A struct {
}

//go:noinline
func (a *A) Foo(s [2]string) {
}
