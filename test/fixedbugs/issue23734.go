// run


package main

func main() {
	m := map[interface{}]int{}
	k := []int{}

	mustPanic(func() {
		_ = m[k]
	})
	mustPanic(func() {
		_, _ = m[k]
	})
	mustPanic(func() {
		delete(m, k)
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
