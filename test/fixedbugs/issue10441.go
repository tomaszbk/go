// compile

package p

func bar() {
	f := func() {}
	foo(&f)
}

//go:noinline
func foo(f *func()) func() {
	return *f
}
