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

var packageCallback func(int) int = (x) => x + packageOffset
var packageOffset = 7

func scenarios() {
	xs := []int{3, 1, 2}
	slices.SortFunc(xs, (a, b) => cmp.Compare(a, b))
	must(fmt.Sprint(xs) == "[1 2 3]", "sort func")
	must(slices.IndexFunc(xs, (x) => x > 1) == 1, "index func")
	sort.Slice(xs, (a, b) => xs[a] > xs[b])
	must(fmt.Sprint(xs) == "[3 2 1]", "sort slice")
	must(strings.Map((r) => r + 1, "abc") == "bcd", "strings map")
	must(fmt.Sprint(Map([]User{{"a"}, {"b"}}, (u) => u.Name)) == "[a b]", "map result inference")
	must(fmt.Sprint(Map([]string{"bad", "7"}, (s) => strconv.Atoi(s) or err {
		return -1
	})) == "[-1 7]", "handler inference")
	must(fmt.Sprint(Map(xs, (x) => 0.5)) == "[0.5 0.5 0.5]", "untyped inference")
	must(Reduce(xs, 0, (acc, x) => acc + x) == 6, "reduce defaults")
	must(Pair(0, 1, (x) => float64(x)) == float64(0), "pair inference")
	must(Pick(() => 1, 2.5) == float64(1), "combine untyped results")
	must(Pick(() => { return 'a' }, 0) == 97, "block defaults")
	once := sync.OnceValues(() => strconv.Atoi("5"))
	n, err := once()
	must(n == 5 && err == nil, "multi-result inference")
	must(Twice(3, (x) => x * 2) == 12, "function constraint")
	must(Map[int, int](xs, (x) => { return x + 1 })[0] == 4, "explicit generic block")
	var f Callback = (x) => x + 1
	must(f(2) == 3 && Callback((x) => x * 2)(3) == 6, "named targets")
	var opts []func(int) int
	opts = append(opts, (x) => x + 2)
	obj := struct{ F func(int) int }{(x) => x + 3}
	callbacks := map[string]func(int) int{"f": (x) => x + 4}
	ch := make(chan func(int) int, 1)
	ch <- (x) => x + 5
	must(opts[0](1) == 3 && obj.F(1) == 4 && callbacks["f"](1) == 5 && (<-ch)(1) == 6, "context targets")

	must(options(1, (x) => x + 1, (x) => x * 2) == 4, "variadic function targets")
	flag := true
	var chosen func(int) int = (if flag { (x) => x + 1 } else { (x) => x - 1 })
	must(chosen(3) == 4, "conditional lambda")
	var curried func(int) func(int) int = (x) => (y) => x + y
	must(curried(2)(3) == 5, "curried")
	var fac func(int) int
	fac = (n) => if n == 0 { 1 } else { n * fac(n-1) }
	must(fac(5) == 120, "recursion")
	captured := 1
	var closure func() int = () => captured
	captured = 9
	must(closure() == 9, "capture by reference")
	var loop []func() int
	for i := range 3 {
		loop = append(loop, () => i)
	}
	must(loop[0]() == 0 && loop[1]() == 1 && loop[2]() == 2, "loop capture")
	trace := ""
	var lazy func() int = () => { trace += "b"; return 1 }
	must(trace == "", "creation lazy")
	must(lazy() == 1 && trace == "b", "invoke lazy")
	trace = ""

	var parse func(string) (int, error) = (s) => { defer func() { trace += "d" }(); v := strconv.Atoi(s)!; return v, nil }
	n, err = parse("bad")
	must(n == 0 && err != nil && trace == "d", "propagation boundary failure")
	n, err = parse("7")
	must(n == 7 && err == nil && trace == "dd", "propagation success")
	var handle func(string) int = (s) => strconv.Atoi(s) or err {
		return -1
	}
	must(handle("bad") == -1 && handle("8") == 8, "handler boundary")
	var noResult func() = () => noop()
	noResult()
	var noErrorResult func() = () => failed() or err {
		trace += "h"
	}
	noErrorResult()
	must(trace == "ddh", "error-only handler")
	var nilFn func() *int = () => nil
	must(nilFn() == nil, "nil result")
	var anyFn func() any = () => 1
	must(anyFn() == 1, "interface result")
	var recoverFn func() = () => { defer func() { must(recover() == "boom", "recover") }(); panic("boom") }
	recoverFn()
	var seq iter.Seq[int] = (yield) => {
		for i := range 3 {
			if !yield(i) {
				return
			}
		}
	}
	var sumSeq func() int = () => {
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
	var early func() (int, error) = () => {
		for range seq {
			strconv.Atoi("bad")!
		}
		return 1, nil
	}
	n, err = early()
	must(n == 0 && err != nil, "range propagation boundary")
	done := make(chan int, 1)
	var goroutine func() = () => { done <- 11 }
	go goroutine()
	must(<-done == 11, "goroutine")
	must(packageCallback(1) == 8, "package initializer")
	fmt.Println("callbacks inference targets closures errors range PASS")
}
