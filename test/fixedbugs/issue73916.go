// run


package main

func callRecover() {
	if recover() != nil {
		println("recovered")
	}
}

func F(int) { callRecover() }

func main() {
	mustPanic(func() {
		defer F(1)
		panic("XXX")
	})
}

func mustPanic(f func()) {
	defer func() {
		r := recover()
		if r == nil {
			panic("didn't panic")
		}
	}()
	f()
}
