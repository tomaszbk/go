package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"unsafe"
)

// Package initialization. The function literal references both branches,
// so initElseValue is initialized before initValue although only the then
// branch runs.
var initValue = func() int {
	if initCond {
		return initThen()
	}
	return initElse()
}()
var initCond = flag("init cond", true)
var initThenValue = side("init then value", 1)
var initElseValue = side("init else value", 2)

func initThen() int { return side("init then", initThenValue*10) }
func initElse() int { return side("init else", initElseValue*10) }

var initLabel = func() string {
	if initCond {
		return "then"
	}
	return "else"
}()
var initWidth = func() string {
	if intSize == 64 {
		return "64-bit"
	}
	return "32-bit"
}()

// Constants emulating a selection, as in math/const.go.
const intSize = 32 << (^uint(0) >> 63)
const debug = false
const mode = "release"
const ratio = 1.0
const width int8 = 2

var buffer [intSize / 8]byte

func label(n int) string {
	var label string
	if n == 1 {
		label = "item"
	} else {
		label = "items"
	}
	return label
}

func lazy(c bool) int {
	var v int
	if flag("cond", c) {
		v = side("then", 1)
	} else {
		v = side("else", 2)
	}
	return v
}

func arguments(c bool) int {
	a := side("a", 1)
	var b int
	if flag("c", c) {
		b = side("b", 2)
	} else {
		b = side("d", 3)
	}
	return sum(a, b, side("e", 4))
}

func propagateArguments(c, fail bool) (int, error) {
	a := side("a", 1)
	var b int
	if flag("c", c) {
		v, err := value("b", 2, fail)
		if err != nil {
			return 0, err
		}
		b = v
	} else {
		b = side("d", 3)
	}
	return sum(a, b, side("e", 4)), nil
}

func guards(s []int, p *point, b []byte) (int, int, bool) {
	var ptr unsafe.Pointer
	if len(b) > 0 {
		ptr = unsafe.Pointer(&b[0])
	}
	first := -1
	if len(s) > 0 {
		first = s[0]
	}
	x := 0
	if p != nil {
		x = p.x
	}
	return first, x, ptr == nil
}

func operators(taxable bool) (int, int, string, bool, int) {
	price, tax := side("price", 100), side("tax", 21)
	var taxed int
	if taxable {
		taxed = tax
	}
	total := price + taxed
	first := taxed + price*2
	name := "net"
	if taxable {
		name += "+tax"
	}
	left := side("left", 1)
	var right int
	if flag("right cond", taxable) {
		right = side("right then", 2)
	} else {
		right = side("right else", 0)
	}
	less := left < right
	neg := -1
	if taxable {
		neg = -tax
	}
	return total, first, name, less, neg
}

func logical(first, c bool) bool {
	if !flag("first", first) {
		return false
	}
	var v bool
	if flag("c", c) {
		v = flag("then", true)
	} else {
		v = flag("else", false)
	}
	return v
}

func propagate(useCache, fail bool) (int, error) {
	var data int
	if useCache {
		v, err := value("cache", 7, fail)
		if err != nil {
			return 0, err
		}
		data = v
	} else {
		v, err := value("fetch", 9, false)
		if err != nil {
			return 0, err
		}
		data = v * 2
	}
	mark("after")
	return data, nil
}

func propagateNamed(c, fail bool) (n int, err error) {
	defer func() { mark(fmt.Sprintf("defer:%d:%v", n, err)) }()
	n = 5
	if c {
		v, callErr := value("named", 8, fail)
		if callErr != nil {
			return 0, callErr
		}
		n = v
	} else {
		n = n + 1
	}
	return n * 10, nil
}

func handled(c, fail bool) (int, error) {
	var v int
	if c {
		n, err := value("primary", 3, fail)
		if err != nil {
			mark("handler")
			return -1, fmt.Errorf("primary: %w", err)
		}
		v = n
	} else {
		v = 4
	}
	return v, nil
}

func iterate(fail bool) (total int, err error) {
	defer func() { mark(fmt.Sprintf("iterate defer:%d:%v", total, err)) }()
	for i := range sequence {
		var v int
		if i == 2 {
			n, callErr := value("two", 20, fail)
			if callErr != nil {
				return 0, callErr
			}
			v = n
		} else {
			v = i
		}
		total += v
	}
	return total, nil
}

func iterateHandled(fail bool) int {
	sum := 0
	for i := range sequence {
		var v int
		if i == 2 {
			n, err := value("two", 20, fail)
			if err != nil {
				mark("loop handler:" + err.Error())
				return -1
			}
			v = n
		} else {
			v = i
		}
		sum += v
	}
	return sum
}

func untypedKinds(c bool) []string {
	var x float64
	var r rune
	var z complex128
	var b byte
	var f32 float32
	if c {
		x, r, z, b, f32 = 1, 'a', 1, 255, 1
	} else {
		x, r, z, b, f32 = 2.5, 1, 2i, 0, 0.5
	}
	return []string{kind(x), kind(r), kind(z), kind(b), kind(f32)}
}

func nilBranches(c bool) []bool {
	n := 1
	var p *int
	var s []int
	var m map[string]int
	var fn func()
	if c {
		p, s, fn = &n, []int{1}, func() {}
	} else {
		m = map[string]int{"k": 1}
	}
	return []bool{p == nil, s == nil, m == nil, fn == nil}
}

func typedNil(failed bool) (bool, string) {
	var p *myErr // nil
	var err error
	if failed {
		err = p
	} else {
		err = nil
	}
	return err == nil, fmt.Sprintf("%T", err)
}

func nilTargets(failed bool) []bool {
	var p *myErr
	var e error
	if failed {
		e = p
	}
	err := e
	h := holder{err: e}
	errs := []error{e}
	m := map[string]error{"k": e}
	ch := make(chan error, 1)
	ch <- e
	return []bool{
		err == nil, h.err == nil, errs[0] == nil, m["k"] == nil, <-ch == nil,
		isNil(e), returned(failed) == nil,
	}
}

func returned(failed bool) error {
	var p *myErr
	if failed {
		return p
	}
	return nil
}

func readers(useString bool) string {
	var r io.Reader
	if useString {
		r = strings.NewReader("string reader")
	} else {
		r = bytes.NewBufferString("buffer")
	}
	data, _ := io.ReadAll(r)
	return string(data)
}

func interfaceUntyped(c bool) string {
	var v float64
	if c {
		v = 0
	} else {
		v = 1.5
	}
	return fmt.Sprintf("%T:%v", v, v)
}

func conversions(c bool) []string {
	i, s := 3, "variable"
	var f float64
	var a any
	var b []byte
	var str string
	if c {
		f, a, b, str = float64(i), i, []byte("bytes"), string('a')
	} else {
		f, a, b, str = 2.5, s, []byte(s), string('b')
	}
	return []string{kind(f), kind(a), string(b), str}
}

func builtinTargets(failed bool) (bool, int, string) {
	var p *myErr
	var e error
	if failed {
		e = p
	}
	errs := append([]error(nil), e)
	m := map[float64]bool{1: true, 2.5: true}
	key := 2.5
	if failed {
		key = 1
	}
	delete(m, key)
	return errs[0] == nil, len(m), panicValue(func() {
		var v any
		if failed {
			v = p
		}
		panic(v)
	})
}

func pick[T any](c bool, a, b T) T {
	if c {
		return a
	}
	return b
}

func orZero[T any](ok bool, v T) T {
	var r T
	if ok {
		r = v
	}
	return r
}

func abs[T number](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func asError[E error](c bool, e E) error {
	if c {
		return e
	}
	return nil
}

func lookup[K comparable, V any](m map[K]V, k K, def V) V {
	if v, ok := m[k]; ok {
		return v
	}
	return def
}

func identity[T any](v T) T { return v }

func genericTargets(failed bool) []bool {
	var p *myErr
	var e error
	if failed {
		e = p
	}
	return []bool{identity[error](e) == nil, identity(e) == nil}
}

func describe[T fmt.Stringer](c bool, a, b T) string {
	r := b
	if c {
		r = a
	}
	return r.String()
}

func constants() []string {
	var release int
	if debug {
		release = side("debug", 1)
	} else {
		release = side("release", 2)
	}
	const size = unsafe.Sizeof(side("never", 1))
	return []string{
		fmt.Sprint(intSize, len(buffer), mode, release, size),
		kind(ratio), kind(width), initWidth,
	}
}

func literals(c bool) ([]string, map[string]int, point, [2]int) {
	answer, key, value, x, first := "no", "off", 0, -1, 4
	if c {
		answer, key, value, x, first = "yes", "on", 1, 1, 3
	}
	s := []string{"first", answer}
	m := map[string]int{key: value}
	p := point{x: x, y: 2}
	a := [...]int{first, 5}
	return s, m, p, a
}

func methods(c bool) (int, int, int) {
	p1, p2 := point{1, 2}, point{3, 4}
	var get func() int
	if c {
		get = p1.sum
	} else {
		get = p2.sum
	}
	p1.x, p2.x = 100, 100
	bound := get()
	var q point
	if c {
		q = p1
	} else {
		q = p2
	}
	called := q.sum()
	var r *point
	if c {
		r = &p1
	} else {
		r = &p2
	}
	return bound, called, r.scaled(2)
}

func controls(c bool) []string {
	var out []string
	limit, step := 3, 2
	if c {
		limit, step = 2, 1
	}
	for i := 0; i < limit; i = i + step {
		out = append(out, fmt.Sprint("for:", i))
	}
	n := 0
	if c {
		n = len(out)
	}
	if n > 0 {
		out = append(out, "if")
	}
	tag := "b"
	if c {
		tag = "a"
	}
	switch tag {
	case "a":
		out = append(out, "switch a")
	case "b":
		out = append(out, "switch b")
	}
	if !c {
		out = append(out, "case")
	}
	str := "xyz"
	if c {
		str = "ab"
	}
	for i, r := range str {
		out = append(out, fmt.Sprint(i, string(r)))
	}
	count := 2
	if c {
		count = 1
	}
	for i := range count {
		out = append(out, fmt.Sprint("range:", i))
	}
	ch := make(chan int, 1)
	sent := 2
	if c {
		sent = 1
	}
	select {
	case ch <- sent:
		out = append(out, fmt.Sprint("send:", <-ch))
	}
	return append(out, fmt.Sprint("len:", len(str)))
}

func indexing(c bool) ([]int, []int, map[string]int, map[string]int, int) {
	s1, s2 := []int{1, 2}, []int{3, 4}
	m1, m2 := map[string]int{}, map[string]int{}
	s, m := s2, m2
	if flag("which", c) {
		s = s1
	}
	if c {
		m = m1
	}
	s[side("index", 0)] = side("value", 9)
	m["k"] += 5
	n := s[1]
	return s1, s2, m1, m2, n
}

func deferred(c bool) {
	defer mark("deferred last")
	var v int
	if flag("defer cond", c) {
		v = side("defer then", 1)
	} else {
		v = side("defer else", 2)
	}
	defer consume(v)
	mark("body")
}

func goroutine(c bool) int {
	ch := make(chan int)
	v := 20
	if c {
		v = 10
	}
	go func(v int) { ch <- v }(v)
	return <-ch
}

func closures(c bool) (int, int) {
	n := 1
	var f func() int
	if c {
		f = func() int { n++; return n }
	} else {
		f = func() int { n += 10; return n }
	}
	g := func() int {
		if c {
			return n * 2
		}
		return n * 3
	}
	return f(), g()
}

func nested(c, d bool) int {
	if c {
		mid := 3
		if d {
			mid = 2
		}
		return sum(1, mid, 4)
	}
	v := 6
	if d {
		v = 5
	}
	return v + side("nested", 1)
}

func values(c bool) ([3]int, [3]int) {
	a1, a2 := [3]int{1, 2, 3}, [3]int{4, 5, 6}
	a := a2
	if c {
		a = a1
	}
	a[0] = 0 // a copy
	if c {
		return a, a1
	}
	return a, a2
}

func statements(n int) string {
	if n == 0 {
		return "zero"
	} else if n == 1 {
		return "one"
	}
	if p := (point{n, n}); p == (point{2, 2}) {
		return "two"
	}
	if n < 0 {
		return "negative"
	}
	return "many"
}
