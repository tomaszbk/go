package main

type problem struct{}

func (*problem) Error() string { return "failure" }

var failure error = &problem{}
var calls int

func read(fail bool) (int, error) {
	calls++
	if fail {
		return 9, failure
	}
	return 42, nil
}
func triple(fail bool) (int, string, error) {
	calls++
	if fail {
		return 9, "partial", failure
	}
	return 42, "ok", nil
}
func only(fail bool) error {
	if fail {
		return failure
	}
	return nil
}
func consume(n int, s string) int {
	if s != "ok" {
		panic(s)
	}
	return n
}
func seq(yield func(int) bool) {
	for i := 0; i < 3; i++ {
		if !yield(i) {
			return
		}
	}
}
func check(ok bool) {
	if !ok {
		panic("assertion")
	}
}

func propagate(fail bool) (n int, err error) {
	n = 99
	defer func() {
		if fail {
			check(n == 0 && err == failure)
			n = 7
		}
	}()
	n = read(fail)!
	return n, nil
}
func local(fail bool) (int, error) {
	n := read(fail) or err {
		return 5, err
	}
	return n, nil
}
func multiple() (int, string) {
	a, b := triple(false) or err {
		panic(err)
	}
	return a, b
}
func callargs() int {
	return consume(triple(false) or err {
		panic(err)
	})
}
func returned() (int, string) {
	return triple(false) or err {
		panic(err)
	}
}
func erroronly() error { only(false)!; only(true)!; panic("unreachable") }
func fallthroughHandler() {
	only(true) or err {
		check(err == failure)
		calls++
	}
}
func short() error {
	if false && read(true)! > 0 {
		panic("unreachable")
	}
	return nil
}
func loop() (int, error) {
	for i := range seq {
		read(i == 1)!
	}
	return 9, nil
}
func aggregate() (s struct{ N int }, a [2]int, err error) { s.N = 9; a[0] = 7; only(true)!; return }
func typednil() error {
	f := func() error { var p *problem; return p }
	f()!
	panic("typed nil was treated as nil")
}
func generic[T any](x T) (T, error) { only(true)!; return x, nil }

func main() {
	n, e := propagate(false)
	check(n == 42 && e == nil)
	n, e = propagate(true)
	check(n == 7 && e == failure)
	n, e = local(true)
	check(n == 5 && e == failure)
	n, e = local(false)
	check(n == 42 && e == nil)
	a, b := multiple()
	check(a == 42 && b == "ok")
	a, b = returned()
	check(a == 42 && b == "ok")
	check(callargs() == 42)
	check(erroronly() == failure)
	fallthroughHandler()
	before := calls
	check(short() == nil)
	check(calls == before)
	n, e = loop()
	check(n == 0 && e == failure)
	s, arr, e := aggregate()
	check(s.N == 0 && arr == [2]int{} && e == failure)
	check(typednil() != nil)
	v, e := generic(123)
	check(v == 0 && e == failure)
	check(calls == 10)
	println("gon analysis pair passed")
}
