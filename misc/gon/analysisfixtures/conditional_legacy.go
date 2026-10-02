package main

// Conditional expressions. conditional_modern.go is the same program with
// conditional expressions; both must print the same output. The programs
// import nothing, so that they also run in the SSA interpreter.

type problem struct{ msg string }

func (p *problem) Error() string { return p.msg }

var failure error = &problem{"failure"}

// trace records the calls that each scenario makes, in order.
var trace string

func note(s string) { trace += s + ";" }

func check(ok bool, what string) {
	if !ok {
		panic("assertion failed: " + what)
	}
}

// expect checks and resets the trace.
func expect(want, what string) {
	check(trace == want, what+": got "+trace+", want "+want)
	trace = ""
}

func flag(name string, v bool) bool { note(name); return v }
func num(name string, v int) int    { note(name); return v }

func read(fail bool) (int, error) {
	note("read")
	if fail {
		return 9, failure
	}
	return 42, nil
}

// Constants. one is a floating-point constant.
const intSize = 32 << (^uint(0) >> 63)
const debug = false
const one = 1.0

var words [intSize / 8]byte

// Package initialization: y depends on pickY, which refers to first and
// second, so second is initialized before y, although only first is used.
var pick = true
var y = pickY()
var first = 1
var second = num("init-second", 2)
var initTrace = trace

func pickY() int {
	if pick {
		return num("init-y", first)
	}
	return second
}

func sum3(a, b, c int) int { return a + b + c }

// order checks f(a(), if c() { b() } else { d() }, e()).
func order(c bool) int {
	a := num("a", 1)
	var mid int
	if flag("c", c) {
		mid = num("b", 10)
	} else {
		mid = num("d", 20)
	}
	return sum3(a, mid, num("e", 100))
}

func propagate(cached, fail bool) (int, error) {
	var v int
	if cached {
		v = num("cache", 7)
	} else {
		r, err := read(fail)
		if err != nil {
			return 0, err
		}
		v = r
	}
	return v + 1, nil
}

func handled(fail bool) (int, error) {
	var v int
	if fail {
		r, err := read(true)
		if err != nil {
			note("handler")
			return -1, err
		}
		v = r
	} else {
		v = num("plain", 5)
	}
	return v, nil
}

func typedNil(failed bool) error {
	var p *problem
	var err error
	if failed {
		err = p
	} else {
		err = nil
	}
	return err
}

type shape interface{ area() int }
type square struct{ s int }
type rect struct{ w, h int }

func (q square) area() int { return q.s * q.s }
func (r *rect) area() int  { return r.w * r.h }

func pickShape(big bool) shape {
	if big {
		return &rect{3, 4}
	}
	return square{2}
}

func choose[T any](c bool, a, b T) T {
	if c {
		return a
	}
	return b
}

func zeroOr[T any](c bool, v T) T {
	var zero T
	if c {
		return v
	}
	return zero
}

func price(taxable bool) (int, int) {
	tax, base := 5, 100
	var added int
	if taxable {
		added = tax
	}
	return base + added, added + base
}

func loops(c bool) (n int) {
	limit := 3
	if c {
		limit = 2
	}
	for i := 0; i < limit; i++ {
		n++
	}
	more := false
	if c {
		more = n > 1
	}
	if more {
		n += 10
	}
	key := "b"
	if c {
		key = "a"
	}
	switch key {
	case "a":
		n += 100
	case "b":
		n += 200
	}
	s := "xyz"
	if c {
		s = "xy"
	}
	for _, r := range s {
		n += int(r - 'x')
	}
	return n
}

func values(c bool) {
	var x float64
	if c {
		x = 1
	} else {
		x = 2.5
	}
	check(x == choose(c, 1.0, 2.5) && x/2 == choose(c, 0.5, 1.25), "untyped branches default to float64")
	str := "cde"
	if c {
		str = "ab"
	}
	check(len(str) == choose(c, 2, 3), "untyped string")
	str = "cd"
	if c {
		str = "ab"
	}
	check(str[1] == choose[byte](c, 'b', 'd'), "index of a string")

	var shift uint = 3
	base := 2
	if c {
		base = 1
	}
	s := base << shift
	check(s == choose(c, 8, 16), "shift")

	elem := 2
	if c {
		elem = 1
	}
	list := []int{elem, 3}
	value := 5
	if c {
		value = 4
	}
	m := map[string]int{"k": value}
	w := 7
	if c {
		w = 6
	}
	r := rect{w: w, h: 1}
	check(list[0] == choose(c, 1, 2) && m["k"] == choose(c, 4, 5) && r.w == choose(c, 6, 7), "composite literals")

	a, b := []int{0}, []int{0}
	target := b
	if c {
		target = a
	}
	target[0] = 9
	check(a[0]+b[0] == 9 && (a[0] == 9) == c, "assignment through an index")
	m1, m2 := map[string]int{}, map[string]int{}
	mt := m2
	if c {
		mt = m1
	}
	mt["k"] = 1
	check(len(m1) == choose(c, 1, 0) && len(m2) == choose(c, 0, 1), "map assignment")

	r1, r2 := &rect{1, 2}, &rect{3, 4}
	rt := r2
	if c {
		rt = r1
	}
	check(rt.w == choose(c, 1, 3), "field selection")
	area := rt.area
	check(area() == choose(c, 2, 12), "method value")

	captured := 1
	var f func() int
	if c {
		f = func() int { return captured }
	} else {
		f = func() int { return -captured }
	}
	captured = 2
	check(f() == choose(c, 2, -2), "closures")

	ch := make(chan int, 1)
	ch <- 8
	got := 0
	if c {
		got = <-ch
	}
	check(got == choose(c, 8, 0) && len(ch) == choose(c, 0, 1), "lazy receive")

	var nested int
	if c {
		inner := 2
		if flag("inner", c) {
			inner = 1
		}
		nested = num("outer", inner)
	} else {
		nested = 3
	}
	check(nested == choose(c, 1, 3), "conditional expression as an operand inside a branch")

	total := 0
	items := []int{4}
	if c {
		items = []int{1, 2}
	}
	for _, v := range items {
		total += v
	}
	check(total == choose(c, 3, 4), "range")
}

func deferred(c bool) {
	arg := "deferred-b"
	if flag("defer-cond", c) {
		arg = "deferred-a"
	}
	defer note(arg)
	note("body")
}

func main() {
	expect("init-second;init-y;", "")
	check(initTrace == "init-second;init-y;" && y == 1, "package initialization")
	check(intSize == 32<<(^uint(0)>>63) && len(words) == intSize/8, "constants")
	check(one/2 == 0.5, "combined constant kind")

	check(order(true) == 111, "order true")
	expect("a;c;b;e;", "order true")
	check(order(false) == 121, "order false")
	expect("a;c;d;e;", "order false")

	n, err := propagate(true, true)
	check(n == 8 && err == nil, "cached")
	expect("cache;", "cached")
	n, err = propagate(false, false)
	check(n == 43 && err == nil, "read")
	expect("read;", "read")
	n, err = propagate(false, true)
	check(n == 0 && err == failure, "propagation")
	expect("read;", "propagation")

	n, err = handled(true)
	check(n == -1 && err == failure, "handler")
	expect("read;handler;", "handler")
	n, err = handled(false)
	check(n == 5 && err == nil, "no handler")
	expect("plain;", "no handler")

	check(typedNil(false) == nil, "nil interface")
	check(typedNil(true) != nil, "typed nil pointer in an interface")

	check(pickShape(true).area() == 12 && pickShape(false).area() == 4, "interface conversion of each branch")
	check(choose(true, "x", "y") == "x" && choose(false, 1, 2) == 2, "generic")
	check(zeroOr(true, "v") == "v" && zeroOr(false, "v") == "" && zeroOr(false, 3) == 0, "generic zero")

	p1, p2 := price(true)
	q1, q2 := price(false)
	check(p1 == 105 && p2 == 105 && q1 == 100 && q2 == 100, "operands of other operations")

	check(loops(true) == 2+10+100+1, "control clauses true")
	check(loops(false) == 3+200+3, "control clauses false")

	values(true)
	expect("inner;outer;", "values true")
	values(false)
	expect("", "values false")

	deferred(true)
	expect("defer-cond;body;deferred-a;", "defer true")
	deferred(false)
	expect("defer-cond;body;deferred-b;", "defer false")

	println("gon conditional pair passed")
}
