package main

// Conditional expressions. conditional_legacy.go is the same program with
// if statements and temporaries; both must print the same output. The
// programs import nothing, so that they also run in the SSA interpreter.

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

// Constants. The untyped branches 1 and 2.5 combine like the operands of a
// binary operation, so one is an untyped floating-point constant.
const intSize = if ^uint(0)>>63 == 1 { 64 } else { 32 }
const debug = false
const one = if !debug { 1 } else { 2.5 }

var words [if intSize == 64 { 8 } else { 4 }]byte

// Package initialization: y depends on the references in both branches,
// so second is initialized before y, although only first is selected.
var pick = true
var y = if pick { num("init-y", first) } else { second }
var first = 1
var second = num("init-second", 2)
var initTrace = trace

func sum3(a, b, c int) int { return a + b + c }

// order checks f(a(), if c() { b() } else { d() }, e()).
func order(c bool) int {
	return sum3(num("a", 1), if flag("c", c) { num("b", 10) } else { num("d", 20) }, num("e", 100))
}

func propagate(cached, fail bool) (int, error) {
	v := if cached { num("cache", 7) } else { read(fail)! }
	return v + 1, nil
}

func handled(fail bool) (int, error) {
	v := if fail {
		read(true) or err {
			note("handler")
			return -1, err
		}
	} else {
		num("plain", 5)
	}
	return v, nil
}

func typedNil(failed bool) error {
	var p *problem
	var err error = if failed { p } else { nil }
	return err
}

type shape interface{ area() int }
type square struct{ s int }
type rect struct{ w, h int }

func (q square) area() int { return q.s * q.s }
func (r *rect) area() int  { return r.w * r.h }

func pickShape(big bool) shape {
	return if big { &rect{3, 4} } else { square{2} }
}

func choose[T any](c bool, a, b T) T { return if c { a } else { b } }

func zeroOr[T any](c bool, v T) T {
	var zero T
	return if c { v } else { zero }
}

func price(taxable bool) (int, int) {
	tax, base := 5, 100
	return base + if taxable { tax } else { 0 }, if taxable { tax } else { 0 } + base
}

func loops(c bool) (n int) {
	for i := 0; i < (if c { 2 } else { 3 }); i++ {
		n++
	}
	if (if c { n > 1 } else { false }) {
		n += 10
	}
	switch if c { "a" } else { "b" } {
	case "a":
		n += 100
	case "b":
		n += 200
	}
	for _, r := range if c { "xy" } else { "xyz" } {
		n += int(r - 'x')
	}
	return n
}

func values(c bool) {
	x := if c { 1 } else { 2.5 }
	check(x == choose(c, 1.0, 2.5) && x/2 == choose(c, 0.5, 1.25), "untyped branches default to float64")
	check(len(if c { "ab" } else { "cde" }) == choose(c, 2, 3), "untyped string")
	check((if c { "ab" } else { "cd" })[1] == choose[byte](c, 'b', 'd'), "index of a string")

	var shift uint = 3
	s := (if c { 1 } else { 2 }) << shift
	check(s == choose(c, 8, 16), "shift")

	list := []int{if c { 1 } else { 2 }, 3}
	m := map[string]int{"k": if c { 4 } else { 5 }}
	r := rect{w: if c { 6 } else { 7 }, h: 1}
	check(list[0] == choose(c, 1, 2) && m["k"] == choose(c, 4, 5) && r.w == choose(c, 6, 7), "composite literals")

	a, b := []int{0}, []int{0}
	(if c { a } else { b })[0] = 9
	check(a[0]+b[0] == 9 && (a[0] == 9) == c, "assignment through an index")
	m1, m2 := map[string]int{}, map[string]int{}
	(if c { m1 } else { m2 })["k"] = 1
	check(len(m1) == choose(c, 1, 0) && len(m2) == choose(c, 0, 1), "map assignment")

	r1, r2 := &rect{1, 2}, &rect{3, 4}
	check((if c { r1 } else { r2 }).w == choose(c, 1, 3), "field selection")
	area := (if c { r1 } else { r2 }).area
	check(area() == choose(c, 2, 12), "method value")

	captured := 1
	f := if c { func() int { return captured } } else { func() int { return -captured } }
	captured = 2
	check(f() == choose(c, 2, -2), "closures")

	ch := make(chan int, 1)
	ch <- 8
	got := if c { <-ch } else { 0 }
	check(got == choose(c, 8, 0) && len(ch) == choose(c, 0, 1), "lazy receive")

	nested := if c { num("outer", if flag("inner", c) { 1 } else { 2 }) } else { 3 }
	check(nested == choose(c, 1, 3), "conditional expression as an operand inside a branch")

	total := 0
	for _, v := range if c { []int{1, 2} } else { []int{4} } {
		total += v
	}
	check(total == choose(c, 3, 4), "range")
}

func deferred(c bool) {
	defer note(if flag("defer-cond", c) { "deferred-a" } else { "deferred-b" })
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
