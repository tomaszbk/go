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
	v, e := read(fail)
	if e != nil {
		return 0, e
	}
	n = v
	return n, nil
}
func local(fail bool) (int, error) {
	n, err := read(fail)
	if err != nil {
		return 5, err
	}
	return n, nil
}
func multiple() (int, string) {
	a, b, err := triple(false)
	if err != nil {
		panic(err)
	}
	return a, b
}
func callargs() int {
	a, b, err := triple(false)
	if err != nil {
		panic(err)
	}
	return consume(a, b)
}
func returned() (int, string) {
	a, b, err := triple(false)
	if err != nil {
		panic(err)
	}
	return a, b
}
func erroronly() error {
	if e := only(false); e != nil {
		return e
	}
	if e := only(true); e != nil {
		return e
	}
	panic("unreachable")
}
func fallthroughHandler() {
	if err := only(true); err != nil {
		check(err == failure)
		calls++
	}
}
func short() error {
	if false {
		v, e := read(true)
		if e != nil {
			return e
		}
		if v > 0 {
			panic("unreachable")
		}
	}
	return nil
}
func loop() (int, error) {
	for i := range seq {
		if _, e := read(i == 1); e != nil {
			return 0, e
		}
	}
	return 9, nil
}
func aggregate() (s struct{ N int }, a [2]int, err error) {
	s.N = 9
	a[0] = 7
	if e := only(true); e != nil {
		return struct{ N int }{}, [2]int{}, e
	}
	return
}
func typednil() error {
	f := func() error { var p *problem; return p }
	if e := f(); e != nil {
		return e
	}
	panic("typed nil was treated as nil")
}
func generic[T any](x T) (T, error) {
	if e := only(true); e != nil {
		var zero T
		return zero, e
	}
	return x, nil
}

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
