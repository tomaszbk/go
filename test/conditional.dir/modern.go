package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"unsafe"
)

// Package initialization. Both branches are dependencies of initValue, so
// initElseValue is initialized before it although only the then branch runs.
var initValue = if initCond { initThen() } else { initElse() }
var initCond = flag("init cond", true)
var initThenValue = side("init then value", 1)
var initElseValue = side("init else value", 2)

func initThen() int { return side("init then", initThenValue*10) }
func initElse() int { return side("init else", initElseValue*10) }

var initLabel = if initCond { "then" } else { "else" }
var initWidth = if intSize == 64 { "64-bit" } else { "32-bit" }

// Constant conditional expressions.
const intSize = if ^uint(0)>>63 == 1 { 64 } else { 32 }
const debug = false
const mode = if debug { "debug" } else { "release" }
const ratio = if debug { 2.5 } else { 1 } // untyped float for either value of debug
const width int8 = if debug { 1 } else { 2 }

var buffer [if intSize == 64 { 8 } else { 4 }]byte

func label(n int) string {
	return if n == 1 { "item" } else { "items" }
}

func lazy(c bool) int {
	return if flag("cond", c) { side("then", 1) } else { side("else", 2) }
}

func arguments(c bool) int {
	return sum(side("a", 1), if flag("c", c) { side("b", 2) } else { side("d", 3) }, side("e", 4))
}

func propagateArguments(c, fail bool) (int, error) {
	return sum(side("a", 1), if flag("c", c) { value("b", 2, fail)! } else { side("d", 3) }, side("e", 4)), nil
}

// The unselected branch is not evaluated, so it can be guarded by the
// condition.
func guards(s []int, p *point, b []byte) (int, int, bool) {
	ptr := if len(b) > 0 { unsafe.Pointer(&b[0]) } else { nil }
	return if len(s) > 0 { s[0] } else { -1 }, if p != nil { p.x } else { 0 }, ptr == nil
}

func operators(taxable bool) (int, int, string, bool, int) {
	price, tax := side("price", 100), side("tax", 21)
	total := price + if taxable { tax } else { 0 }
	first := if taxable { tax } else { 0 } + price*2
	name := "net" + if taxable { "+tax" } else { "" }
	less := side("left", 1) < if flag("right cond", taxable) { side("right then", 2) } else { side("right else", 0) }
	neg := -if taxable { tax } else { 1 }
	return total, first, name, less, neg
}

func logical(first, c bool) bool {
	return flag("first", first) && if flag("c", c) { flag("then", true) } else { flag("else", false) }
}

func propagate(useCache, fail bool) (int, error) {
	data := if useCache { value("cache", 7, fail)! } else { value("fetch", 9, false)! * 2 }
	mark("after")
	return data, nil
}

func propagateNamed(c, fail bool) (n int, err error) {
	defer func() { mark(fmt.Sprintf("defer:%d:%v", n, err)) }()
	n = 5
	n = if c { value("named", 8, fail)! } else { n + 1 }
	return n * 10, nil
}

func handled(c, fail bool) (int, error) {
	v := if c {
		value("primary", 3, fail) or err {
			mark("handler")
			return -1, fmt.Errorf("primary: %w", err)
		}
	} else {
		4
	}
	return v, nil
}

func iterate(fail bool) (total int, err error) {
	defer func() { mark(fmt.Sprintf("iterate defer:%d:%v", total, err)) }()
	for i := range sequence {
		total += if i == 2 { value("two", 20, fail)! } else { i }
	}
	return total, nil
}

func iterateHandled(fail bool) int {
	sum := 0
	for i := range sequence {
		sum += if i == 2 {
			value("two", 20, fail) or err {
				mark("loop handler:" + err.Error())
				return -1
			}
		} else {
			i
		}
	}
	return sum
}

func untypedKinds(c bool) []string {
	x := if c { 1 } else { 2.5 }
	r := if c { 'a' } else { 1 }
	z := if c { 1 } else { 2i }
	var b byte = if c { 255 } else { 0 }
	var f32 float32 = if c { 1 } else { 0.5 }
	return []string{kind(x), kind(r), kind(z), kind(b), kind(f32)}
}

func nilBranches(c bool) []bool {
	n := 1
	p := if c { &n } else { nil }
	s := if c { []int{1} } else { nil }
	var m map[string]int = if c { nil } else { map[string]int{"k": 1} }
	fn := if c { func() {} } else { nil }
	return []bool{p == nil, s == nil, m == nil, fn == nil}
}

func typedNil(failed bool) (bool, string) {
	var p *myErr // nil
	var err error = if failed { p } else { nil }
	return err == nil, fmt.Sprintf("%T", err)
}

// Each branch is converted to the target type separately, so the nil
// branch is a nil error in every target context.
func nilTargets(failed bool) []bool {
	var p *myErr
	var err error
	err = if failed { p } else { nil }
	h := holder{err: if failed { p } else { nil }}
	errs := []error{if failed { p } else { nil }}
	m := map[string]error{"k": if failed { p } else { nil }}
	ch := make(chan error, 1)
	ch <- if failed { p } else { nil }
	return []bool{
		err == nil, h.err == nil, errs[0] == nil, m["k"] == nil, <-ch == nil,
		isNil(if failed { p } else { nil }), returned(failed) == nil,
	}
}

func returned(failed bool) error {
	var p *myErr
	return if failed { p } else { nil }
}

func readers(useString bool) string {
	var r io.Reader = if useString { strings.NewReader("string reader") } else { bytes.NewBufferString("buffer") }
	data, _ := io.ReadAll(r)
	return string(data)
}

// With an interface target, untyped constant branches first get the type
// they have without a target: here the combined kind is floating-point.
func interfaceUntyped(c bool) string {
	return fmt.Sprintf("%T:%v", if c { 0 } else { 1.5 }, if c { 0 } else { 1.5 })
}

// A conversion applies to each branch separately.
func conversions(c bool) []string {
	i, s := 3, "variable"
	return []string{
		kind(float64(if c { i } else { 2.5 })),
		kind(any(if c { i } else { s })),
		string([]byte(if c { "bytes" } else { s })),
		string(if c { 'a' } else { 'b' }),
	}
}

// Element arguments of append, delete keys and panic arguments are targets.
func builtinTargets(failed bool) (bool, int, string) {
	var p *myErr
	errs := append([]error(nil), if failed { p } else { nil })
	m := map[float64]bool{1: true, 2.5: true}
	delete(m, if failed { 1 } else { 2.5 })
	return errs[0] == nil, len(m), panicValue(func() { panic(if failed { p } else { nil }) })
}

func pick[T any](c bool, a, b T) T { return if c { a } else { b } }

func orZero[T any](ok bool, v T) T {
	var zero T
	return if ok { v } else { zero }
}

func abs[T number](x T) T { return if x < 0 { -x } else { x } }

func asError[E error](c bool, e E) error { return if c { e } else { nil } }

func lookup[K comparable, V any](m map[K]V, k K, def V) V {
	v, ok := m[k]
	return if ok { v } else { def }
}

func identity[T any](v T) T { return v }

// A fully instantiated generic parameter is a target; a conversion gives the
// branches a target in a generic call.
func genericTargets(failed bool) []bool {
	var p *myErr
	return []bool{
		identity[error](if failed { p } else { nil }) == nil,
		identity(error(if failed { p } else { nil })) == nil,
	}
}

func describe[T fmt.Stringer](c bool, a, b T) string { return (if c { a } else { b }).String() }

func constants() []string {
	release := if debug { side("debug", 1) } else { side("release", 2) }
	const size = unsafe.Sizeof(if flag("never", true) { side("never", 1) } else { 2 })
	return []string{
		fmt.Sprint(intSize, len(buffer), mode, release, size),
		kind(ratio), kind(width), initWidth,
	}
}

func literals(c bool) ([]string, map[string]int, point, [2]int) {
	s := []string{"first", if c { "yes" } else { "no" }}
	m := map[string]int{if c { "on" } else { "off" }: if c { 1 } else { 0 }}
	p := point{x: if c { 1 } else { -1 }, y: 2}
	a := [...]int{if c { 3 } else { 4 }, 5}
	return s, m, p, a
}

func methods(c bool) (int, int, int) {
	p1, p2 := point{1, 2}, point{3, 4}
	get := (if c { p1 } else { p2 }).sum
	p1.x, p2.x = 100, 100
	return get(), (if c { p1 } else { p2 }).sum(), (if c { &p1 } else { &p2 }).scaled(2)
}

func controls(c bool) []string {
	var out []string
	for i := 0; i < if c { 2 } else { 3 }; i = i + if c { 1 } else { 2 } {
		out = append(out, fmt.Sprint("for:", i))
	}
	if (if c { len(out) } else { 0 }) > 0 {
		out = append(out, "if")
	}
	switch if c { "a" } else { "b" } {
	case "a":
		out = append(out, "switch a")
	case "b":
		out = append(out, "switch b")
	}
	switch {
	case if c { false } else { true }:
		out = append(out, "case")
	}
	for i, r := range if c { "ab" } else { "xyz" } {
		out = append(out, fmt.Sprint(i, string(r)))
	}
	for i := range if c { 1 } else { 2 } {
		out = append(out, fmt.Sprint("range:", i))
	}
	ch := make(chan int, 1)
	select {
	case ch <- if c { 1 } else { 2 }:
		out = append(out, fmt.Sprint("send:", <-ch))
	}
	return append(out, fmt.Sprint("len:", len(if c { "ab" } else { "xyz" })))
}

func indexing(c bool) ([]int, []int, map[string]int, map[string]int, int) {
	s1, s2 := []int{1, 2}, []int{3, 4}
	m1, m2 := map[string]int{}, map[string]int{}
	(if flag("which", c) { s1 } else { s2 })[side("index", 0)] = side("value", 9)
	(if c { m1 } else { m2 })["k"] += 5
	n := (if c { s1 } else { s2 })[1]
	return s1, s2, m1, m2, n
}

func deferred(c bool) {
	defer mark("deferred last")
	defer consume(if flag("defer cond", c) { side("defer then", 1) } else { side("defer else", 2) })
	mark("body")
}

func goroutine(c bool) int {
	ch := make(chan int)
	go func(v int) { ch <- v }(if c { 10 } else { 20 })
	return <-ch
}

func closures(c bool) (int, int) {
	n := 1
	f := if c { func() int { n++; return n } } else { func() int { n += 10; return n } }
	g := func() int { return if c { n * 2 } else { n * 3 } }
	return f(), g()
}

func nested(c, d bool) int {
	return if c { sum(1, if d { 2 } else { 3 }, 4) } else { (if d { 5 } else { 6 }) + side("nested", 1) }
}

func values(c bool) ([3]int, [3]int) {
	a1, a2 := [3]int{1, 2, 3}, [3]int{4, 5, 6}
	a := if c { a1 } else { a2 }
	a[0] = 0 // the result is a copy
	return a, if c { a1 } else { a2 }
}

func statements(n int) string {
	if n == 0 {
		return "zero"
	} else if n == 1 {
		return if n > 0 { "one" } else { "impossible" }
	}
	if p := (point{n, n}); p == (point{2, 2}) {
		return "two"
	}
	return if n < 0 {
		"negative"
	} else {
		"many"
	}
}
