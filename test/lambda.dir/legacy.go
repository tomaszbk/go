package main

import (
	"cmp"
	"fmt"
	"iter"
	"slices"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var packageCallback func(int) int = func(x int) int { return x + packageOffset }
var packageOffset = 7

func scenarios() {
	xs := []int{3, 1, 2}
	slices.SortFunc(xs, func(a, b int) int { return cmp.Compare(a, b) })
	must(fmt.Sprint(xs) == "[1 2 3]", "sort func")
	must(slices.IndexFunc(xs, func(x int) bool { return x > 1 }) == 1, "index func")
	sort.Slice(xs, func(a, b int) bool { return xs[a] > xs[b] })
	must(fmt.Sprint(xs) == "[3 2 1]", "sort slice")
	must(strings.Map(func(r rune) rune { return r + 1 }, "abc") == "bcd", "strings map")
	must(fmt.Sprint(Map([]User{{"a"}, {"b"}}, func(u User) string { return u.Name })) == "[a b]", "map result inference")
	must(fmt.Sprint(Map([]string{"bad", "7"}, func(s string) int {
		n, err := strconv.Atoi(s)
		if err != nil {
			return -1
		}
		return n
	})) == "[-1 7]", "handler inference")
	must(fmt.Sprint(Map(xs, func(x int) float64 { return 0.5 })) == "[0.5 0.5 0.5]", "untyped inference")
	must(Reduce(xs, 0, func(acc, x int) int { return acc + x }) == 6, "reduce defaults")
	must(Pair(0, 1, func(x int) float64 { return float64(x) }) == float64(0), "pair inference")
	must(Pick(func() float64 { return 1 }, 2.5) == float64(1), "combine untyped results")
	must(Pick(func() int { return 'a' }, 0) == 97, "block defaults")
	once := sync.OnceValues(func() (int, error) { return strconv.Atoi("5") })
	n, err := once()
	must(n == 5 && err == nil, "multi-result inference")
	must(Twice(3, func(x int) int { return x * 2 }) == 12, "function constraint")
	must(Map[int, int](xs, func(x int) int { return x + 1 })[0] == 4, "explicit generic block")
	var f Callback = func(x int) int { return x + 1 }
	must(f(2) == 3 && Callback(func(x int) int { return x * 2 })(3) == 6, "named targets")
	var opts []func(int) int
	opts = append(opts, func(x int) int { return x + 2 })
	obj := struct{ F func(int) int }{func(x int) int { return x + 3 }}
	callbacks := map[string]func(int) int{"f": func(x int) int { return x + 4 }}
	ch := make(chan func(int) int, 1)
	ch <- func(x int) int { return x + 5 }
	must(opts[0](1) == 3 && obj.F(1) == 4 && callbacks["f"](1) == 5 && (<-ch)(1) == 6, "context targets")

	must(options(1, func(x int) int { return x + 1 }, func(x int) int { return x * 2 }) == 4, "variadic function targets")
	flag := true
	var chosen func(int) int
	if flag {
		chosen = func(x int) int { return x + 1 }
	} else {
		chosen = func(x int) int { return x - 1 }
	}
	must(chosen(3) == 4, "conditional lambda")
	var curried func(int) func(int) int = func(x int) func(int) int { return func(y int) int { return x + y } }
	must(curried(2)(3) == 5, "curried")
	var fac func(int) int
	fac = func(n int) int {
		if n == 0 {
			return 1
		}
		return n * fac(n-1)
	}
	must(fac(5) == 120, "recursion")
	captured := 1
	var closure func() int = func() int { return captured }
	captured = 9
	must(closure() == 9, "capture by reference")
	var loop []func() int
	for i := range 3 {
		loop = append(loop, func() int { return i })
	}
	must(loop[0]() == 0 && loop[1]() == 1 && loop[2]() == 2, "loop capture")
	trace := ""
	var lazy func() int = func() int { trace += "b"; return 1 }
	must(trace == "", "creation lazy")
	must(lazy() == 1 && trace == "b", "invoke lazy")
	trace = ""

	var parse func(string) (int, error) = func(s string) (int, error) {
		defer func() { trace += "d" }()
		v, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}
		return v, nil
	}
	n, err = parse("bad")
	must(n == 0 && err != nil && trace == "d", "propagation boundary failure")
	n, err = parse("7")
	must(n == 7 && err == nil && trace == "dd", "propagation success")
	var handle func(string) int = func(s string) int {
		v, err := strconv.Atoi(s)
		if err != nil {
			return -1
		}
		return v
	}
	must(handle("bad") == -1 && handle("8") == 8, "handler boundary")
	var noResult func() = func() { noop() }
	noResult()
	var noErrorResult func() = func() {
		if err := failed(); err != nil {
			trace += "h"
		}
	}
	noErrorResult()
	must(trace == "ddh", "error-only handler")
	var nilFn func() *int = func() *int { return nil }
	must(nilFn() == nil, "nil result")
	var anyFn func() any = func() any { return 1 }
	must(anyFn() == 1, "interface result")
	var recoverFn func() = func() { defer func() { must(recover() == "boom", "recover") }(); panic("boom") }
	recoverFn()
	var seq iter.Seq[int] = func(yield func(int) bool) {
		for i := range 3 {
			if !yield(i) {
				return
			}
		}
	}
	var sumSeq func() int = func() int {
		total := 0
		for i := range seq {
			total += i
		}
		return total
	}
	must(sumSeq() == 3, "range in lambda")
	sum := 0
	for i := range seq {
		sum += i
	}
	must(sum == 3, "lambda iterator")
	var early func() (int, error) = func() (int, error) {
		for range seq {
			_, err := strconv.Atoi("bad")
			if err != nil {
				return 0, err
			}
		}
		return 1, nil
	}
	n, err = early()
	must(n == 0 && err != nil, "range propagation boundary")
	done := make(chan int, 1)
	var goroutine func() = func() { done <- 11 }
	go goroutine()
	must(<-done == 11, "goroutine")
	must(packageCallback(1) == 8, "package initializer")
	fmt.Println("callbacks inference targets closures errors range PASS")
}
