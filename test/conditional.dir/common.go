package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

var trace []string

func mark(s string) { trace = append(trace, s) }

func side(label string, n int) int { mark(label); return n }

func flag(label string, b bool) bool { mark(label); return b }

func sum(a, b, c int) int { mark("sum"); return a + b + c }

func consume(n int) { mark(fmt.Sprint("consume:", n)) }

var failure = errors.New("failure")

func value(label string, n int, fail bool) (int, error) {
	mark(label)
	if fail {
		// A failed conventional call can still have a useful partial result.
		return n, failure
	}
	return n, nil
}

type myErr struct{}

func (*myErr) Error() string { return "typed nil" }

type holder struct{ err error }

func isNil(err error) bool { return err == nil }

type point struct{ x, y int }

func (p point) sum() int { return p.x + p.y }

func (p point) String() string { return fmt.Sprintf("(%d,%d)", p.x, p.y) }

func (p *point) scaled(k int) int {
	p.x *= k
	p.y *= k
	return p.sum()
}

type number interface{ ~int | ~float64 }

func sequence(yield func(int) bool) {
	mark("iterator begin")
	for i := 1; i <= 3; i++ {
		if !yield(i) {
			mark("iterator stopped")
			return
		}
	}
	mark("iterator end")
}

func kind(v any) string { return fmt.Sprintf("%T:%v", v, v) }

func panicValue(f func()) (s string) {
	defer func() { s = fmt.Sprintf("%T", recover()) }()
	f()
	return "no panic"
}

func check(label string, got, want any) {
	if !reflect.DeepEqual(got, want) {
		panic(fmt.Sprintf("%s: got %#v, want %#v", label, got, want))
	}
	fmt.Printf("%s: %v\n", label, got)
}

func checkTrace(label string, want ...string) {
	check(label+" trace", trace, want)
	trace = nil
}

// Ordinary if statements keep their meaning next to conditional expressions.
func compatibility(n int) []string {
	var out []string
	if n > 0 {
		out = append(out, "positive")
	} else if n < 0 {
		out = append(out, "negative")
	} else {
		out = append(out, "zero")
	}
	if p := (point{}); p == (point{}) {
		out = append(out, "composite condition")
	}
	if v := n * 2; v > 2 {
		out = append(out, "init statement")
	}
	for _, p := range []point{{1, 2}} {
		if p == (point{1, 2}) {
			out = append(out, "loop literal")
		}
	}
	return out
}

func main() {
	// Package initialization ran before main.
	checkTrace("init", "init cond", "init then value", "init else value", "init then")
	check("init values", []any{initValue, initLabel, initWidth}, []any{10, "then", strconv.Itoa(strconv.IntSize) + "-bit"})

	check("label", []string{label(1), label(2)}, []string{"item", "items"})

	first, x, isNil := guards(nil, nil, nil)
	check("guards unselected", []any{first, x, isNil}, []any{-1, 0, true})
	first, x, isNil = guards([]int{5}, &point{7, 8}, []byte{1})
	check("guards selected", []any{first, x, isNil}, []any{5, 7, false})

	for _, c := range []bool{true, false} {
		check("lazy", lazy(c), map[bool]int{true: 1, false: 2}[c])
		if c {
			checkTrace("lazy", "cond", "then")
		} else {
			checkTrace("lazy", "cond", "else")
		}

		check("arguments", arguments(c), map[bool]int{true: 7, false: 8}[c])
		if c {
			checkTrace("arguments", "a", "c", "b", "e", "sum")
		} else {
			checkTrace("arguments", "a", "c", "d", "e", "sum")
		}

		total, first, name, less, neg := operators(c)
		if c {
			check("operators", []any{total, first, name, less, neg}, []any{121, 221, "net+tax", true, -21})
			checkTrace("operators", "price", "tax", "left", "right cond", "right then")
		} else {
			check("operators", []any{total, first, name, less, neg}, []any{100, 200, "net", false, -1})
			checkTrace("operators", "price", "tax", "left", "right cond", "right else")
		}

		for _, fail := range []bool{false, true} {
			v, err := propagate(c, fail)
			switch {
			case c && fail:
				check("propagate failure", []any{v, err}, []any{0, failure})
				checkTrace("propagate", "cache")
			case c:
				check("propagate cache", []any{v, err}, []any{7, nil})
				checkTrace("propagate", "cache", "after")
			default:
				check("propagate fetch", []any{v, err}, []any{18, nil})
				checkTrace("propagate", "fetch", "after")
			}

			v, err = propagateArguments(c, fail)
			switch {
			case c && fail:
				check("arguments failure", []any{v, err}, []any{0, failure})
				checkTrace("propagate arguments", "a", "c", "b")
			case c:
				check("arguments success", []any{v, err}, []any{7, nil})
				checkTrace("propagate arguments", "a", "c", "b", "e", "sum")
			default:
				check("arguments else", []any{v, err}, []any{8, nil})
				checkTrace("propagate arguments", "a", "c", "d", "e", "sum")
			}

			v, err = propagateNamed(c, fail)
			switch {
			case c && fail:
				check("named failure", []any{v, err}, []any{0, failure})
				checkTrace("named", "named", "defer:0:failure")
			case c:
				check("named success", []any{v, err}, []any{80, nil})
				checkTrace("named", "named", "defer:80:<nil>")
			default:
				check("named else", []any{v, err}, []any{60, nil})
				checkTrace("named", "defer:60:<nil>")
			}

			v, err = handled(c, fail)
			switch {
			case c && fail:
				check("handled failure", []any{v, errors.Is(err, failure), err.Error()}, []any{-1, true, "primary: failure"})
				checkTrace("handled", "primary", "handler")
			case c:
				check("handled success", []any{v, err}, []any{3, nil})
				checkTrace("handled", "primary")
			default:
				check("handled else", []any{v, err}, []any{4, nil})
				checkTrace("handled")
			}
		}

		check("untyped kinds", untypedKinds(c), map[bool][]string{
			true:  {"float64:1", "int32:97", "complex128:(1+0i)", "uint8:255", "float32:1"},
			false: {"float64:2.5", "int32:1", "complex128:(0+2i)", "uint8:0", "float32:0.5"},
		}[c])
		check("nil branches", nilBranches(c), []bool{!c, !c, c, !c})
		isNilErr, typ := typedNil(c)
		check("typed nil", []any{isNilErr, typ}, map[bool][]any{true: {false, "*main.myErr"}, false: {true, "<nil>"}}[c])
		check("nil targets", nilTargets(c), []bool{!c, !c, !c, !c, !c, !c, !c})
		check("interface target", readers(c), map[bool]string{true: "string reader", false: "buffer"}[c])
		check("interface untyped", interfaceUntyped(c), map[bool]string{true: "float64:0", false: "float64:1.5"}[c])
		check("conversions", conversions(c), map[bool][]string{
			true:  {"float64:3", "int:3", "bytes", "a"},
			false: {"float64:2.5", "string:variable", "variable", "b"},
		}[c])
		appendNil, mapLen, panicked := builtinTargets(c)
		check("builtin targets", []any{appendNil, mapLen, panicked}, map[bool][]any{
			true:  {false, 1, "*main.myErr"},
			false: {true, 1, "*runtime.PanicNilError"},
		}[c])

		x, y := 1, 2
		check("generic pick", []any{pick(c, 1, 2), pick(c, "a", "b"), *pick(c, &x, &y), pick[any](c, 1, "b"), pick(c, point{1, 2}, point{3, 4})},
			map[bool][]any{true: {1, "a", 1, 1, point{1, 2}}, false: {2, "b", 2, "b", point{3, 4}}}[c])
		check("generic zero", []any{orZero(c, []int{1}) == nil, orZero(c, "v"), orZero[error](c, failure)},
			map[bool][]any{true: {false, "v", failure}, false: {true, "", nil}}[c])
		check("generic typed nil", asError(c, (*myErr)(nil)) == nil, !c)
		check("generic targets", genericTargets(c), []bool{!c, !c})

		check("generic lookup", lookup(map[string]int{"k": 1}, map[bool]string{true: "k", false: "missing"}[c], -1), map[bool]int{true: 1, false: -1}[c])
		check("generic method", describe(c, point{1, 2}, point{3, 4}), map[bool]string{true: "(1,2)", false: "(3,4)"}[c])

		s, m, p, a := literals(c)
		if c {
			check("literals", []any{s, m, p, a}, []any{[]string{"first", "yes"}, map[string]int{"on": 1}, point{1, 2}, [2]int{3, 5}})
		} else {
			check("literals", []any{s, m, p, a}, []any{[]string{"first", "no"}, map[string]int{"off": 0}, point{-1, 2}, [2]int{4, 5}})
		}

		bound, called, scaled := methods(c)
		check("methods", []int{bound, called, scaled}, map[bool][]int{true: {3, 102, 204}, false: {7, 104, 208}}[c])

		check("controls", controls(c), map[bool][]string{
			true:  {"for:0", "for:1", "if", "switch a", "0a", "1b", "range:0", "send:1", "len:2"},
			false: {"for:0", "for:2", "switch b", "case", "0x", "1y", "2z", "range:0", "range:1", "send:2", "len:3"},
		}[c])

		s1, s2, m1, m2, n := indexing(c)
		if c {
			check("indexing", []any{s1, s2, m1, m2, n}, []any{[]int{9, 2}, []int{3, 4}, map[string]int{"k": 5}, map[string]int{}, 2})
		} else {
			check("indexing", []any{s1, s2, m1, m2, n}, []any{[]int{1, 2}, []int{9, 4}, map[string]int{}, map[string]int{"k": 5}, 4})
		}
		checkTrace("indexing", "which", "index", "value")

		deferred(c)
		if c {
			checkTrace("deferred", "defer cond", "defer then", "body", "consume:1", "deferred last")
		} else {
			checkTrace("deferred", "defer cond", "defer else", "body", "consume:2", "deferred last")
		}
		check("goroutine", goroutine(c), map[bool]int{true: 10, false: 20}[c])

		f, g := closures(c)
		check("closures", []int{f, g}, map[bool][]int{true: {2, 4}, false: {11, 33}}[c])

		for _, d := range []bool{true, false} {
			want := map[[2]bool]int{{true, true}: 7, {true, false}: 8, {false, true}: 6, {false, false}: 7}[[2]bool{c, d}]
			check("nested", nested(c, d), want)
			if c {
				checkTrace("nested", "sum")
			} else {
				checkTrace("nested", "nested")
			}
		}

		copied, original := values(c)
		check("values", []any{copied, original}, map[bool][]any{true: {[3]int{0, 2, 3}, [3]int{1, 2, 3}}, false: {[3]int{0, 5, 6}, [3]int{4, 5, 6}}}[c])

		for _, first := range []bool{true, false} {
			got := logical(first, c)
			switch {
			case !first:
				check("logical skip", got, false)
				checkTrace("logical", "first")
			case c:
				check("logical then", got, true)
				checkTrace("logical", "first", "c", "then")
			default:
				check("logical else", got, false)
				checkTrace("logical", "first", "c", "else")
			}
		}
	}

	for _, fail := range []bool{false, true} {
		total, err := iterate(fail)
		if fail {
			check("iterate failure", []any{total, err}, []any{0, failure})
			checkTrace("iterate", "iterator begin", "two", "iterator stopped", "iterate defer:0:failure")
		} else {
			check("iterate success", []any{total, err}, []any{24, nil})
			checkTrace("iterate", "iterator begin", "two", "iterator end", "iterate defer:24:<nil>")
		}
		if fail {
			check("iterate handled failure", iterateHandled(fail), -1)
			checkTrace("iterate handled", "iterator begin", "two", "loop handler:failure", "iterator stopped")
		} else {
			check("iterate handled success", iterateHandled(fail), 24)
			checkTrace("iterate handled", "iterator begin", "two", "iterator end")
		}
	}

	check("constants", constants(), []string{
		fmt.Sprint(strconv.IntSize, strconv.IntSize/8, "release", 2, unsafe.Sizeof(0)),
		"float64:1", "int8:2", strconv.Itoa(strconv.IntSize) + "-bit",
	})
	checkTrace("constants", "release")

	check("statements", []string{statements(0), statements(1), statements(2), statements(-1), statements(5)},
		[]string{"zero", "one", "two", "negative", "many"})
	check("compatibility", [][]string{compatibility(2), compatibility(-1), compatibility(0)}, [][]string{
		{"positive", "composite condition", "init statement", "loop literal"},
		{"negative", "composite condition", "loop literal"},
		{"zero", "composite condition", "loop literal"},
	})
	checkTrace("end")
}
