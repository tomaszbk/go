package main

import (
	"fmt"
	"unsafe"
)

func choose[T ~*int](p, q T) T { return p ?? q }

var initial *int = (*int)(nil) ?? &one

func scenarios() {
	leaf := &Node{P: &one, Count: 7, Items: []*int{&one}, Value: &one, Fn: func(x int) *int { return &one }}
	for _, n := range []*Node{nil, {Next: nil}, {Next: leaf}} {
		trace = ""
		p := root(n)?.Next?.P
		want := (*int)(nil)
		if n != nil && n.Next != nil {
			want = n.Next.P
		}
		must(p == want && trace == "r", "chain guards")
		count := n?.Next?.Count ?? 4
		wantCount := 4
		if n != nil && n.Next != nil {
			wantCount = n.Next.Count
		}
		must(count == wantCount, "scalar absence")
		trace = ""
		p = collect(first(), root(n)?.Next?.Get(arg()) ?? fallback(), last())
		wantTrace := "1rf2"
		if n != nil && n.Next != nil {
			wantTrace = "1ram2"
		}
		must(trace == wantTrace, "lexical order "+trace)
		must(*p == 2 || *p == 1, "call value")
	}
	must(initial == &one, "package init")
	trace = ""
	var f func(int) *int
	must(f?(arg()) == nil && trace == "", "nil safe call")
	f = leaf.Fn
	must(f?(arg()) == &one && trace == "a", "safe function call")
	var n *Node
	trace = ""
	must(n?.NilMethod() == nil && trace == "", "skip nil receiver")
	must(n.NilMethod() == &two && trace == "n", "ordinary nil receiver")
	var i Reader = n
	trace = ""
	must(i?.NilMethod() == &two && trace == "n", "typed nil interface")
	i = nil
	trace = ""
	must(i?.NilMethod() == nil && trace == "", "nil interface")
	method := n?.NilMethod
	must(method == nil, "nil method value")
	method = leaf?.NilMethod
	must(method() == &one, "present method value")
	must(n?.ByValue() == nil, "value receiver nil guard")
	must(panics(func() { _ = (&Node{})?.Next.P }), "unguarded tail panics")
	must(panics(func() { _ = (n?.Next).P }), "parenthesized boundary")
	must(panics(func() { _ = (&Embedded{})?.P }), "promoted embedded pointer panics")
	must(leaf?.Items[0] == &one, "guard index")
	must(n?.Items[9] == nil, "skip index")
	must(leaf?.Items[:1][0] == &one, "guard slice")
	must(leaf?.Value.(*int) == &one, "guard assertion")
	must(n?.Value.(*int) == nil, "skip assertion")
	var e *Errors
	var err error = e?.E
	must(err == nil, "absent interface target")
	e = &Errors{}
	err = e?.E
	must(err != nil, "present typed nil interface target")
	var chosenErr error = e?.E ?? sentinel
	must(chosenErr == sentinel, "chain nil test before conversion")
	trace = ""
	x := leaf?.Count ?? numberDefault()
	must(x == 7 && trace == "", "skip scalar fallback")
	x = n?.Count ?? numberDefault()
	must(x == 9 && trace == "d", "scalar fallback")
	var ptr *int
	must((*ptr ?? 10) == 10, "guarded deref")
	ptr = &one
	must((*ptr ?? 10) == 1, "present deref")
	var pp **int
	must((**pp ?? 10) == 10, "nested deref")
	pp = &ptr
	must((**pp ?? 10) == 1, "present nested deref")
	must((*n?.P ?? 10) == 10, "deref chain")
	var pe *error
	must((*pe ?? sentinel) == sentinel, "deref nil interface")
	var emptyError error
	pe = &emptyError
	must((*pe ?? sentinel) == sentinel, "deref nil value")
	var s []int
	must(len(s ?? []int{1}) == 1, "slice nil")
	s = []int{}
	must(len(s ?? []int{1}) == 0, "slice empty")
	var m map[int]int
	must(len(m ?? map[int]int{1: 1}) == 1, "map nil")
	m = map[int]int{}
	must(len(m ?? map[int]int{1: 1}) == 0, "map empty")
	var ch chan int
	ch2 := make(chan int)
	must((ch ?? ch2) == ch2, "channel")
	var fn func() int
	fn = fn ?? () => 4
	must(fn() == 4, "function")
	var u unsafe.Pointer
	must((u ?? unsafe.Pointer(&one)) == unsafe.Pointer(&one), "unsafe pointer")
	var anyNil any = (*int)(nil)
	must((anyNil ?? any(&two)) == anyNil, "interface nil test")
	ptr = nil
	var boxed any = ptr ?? &two
	must(boxed == &two, "nil test before target conversion")
	trace = ""
	must((ptr ?? fallback() ?? &one) == &two && trace == "f", "right associative")
	must(choose((*int)(nil), &one) == &one, "generic")
	must(len(chooseKinds([]int(nil), []int{1})) == 1 && len(chooseKinds(map[int]int(nil), map[int]int{1: 1})) == 1, "mixed nilable type set")
	must(chooseAsAny((*int)(nil), &one) == &one, "generic target conversion")
	trace = ""
	slots := map[int]*int{}
	slots[key()] ??= fallback()
	must(slots[0] == &two && trace == "kf", "map coalescing assignment")
	trace = ""
	slots[key()] ??= fallback()
	must(trace == "k", "no redundant store")
	trace = ""
	getNode(leaf).P ??= fallback()
	must(trace == "g", "field assignment operands")
	leaf.P = nil
	trace = ""
	getNode(leaf).P ??= fallback()
	must(trace == "gf" && leaf.P == &two, "field lazy assignment")
	var local *int
	local ??= &one
	local ??= &two
	must(local == &one, "variable assignment")
	trace = ""
	p, er := load(n, true)
	must(p == nil && er == nil && trace == "", "nil skips propagation")
	trace = ""
	p, er = load(leaf, true)
	must(p == nil && er == sentinel && trace == "l", "failure propagation")
	trace = ""
	p, er = load(leaf, false)
	must(p == &two && er == nil && trace == "l", "successful propagation")
	trace = ""
	p = handle(n, true)
	must(p == nil && trace == "", "nil skips handler")
	trace = ""
	p = handle(leaf, true)
	must(p == &one && trace == "lh", "handler branch")
	trace = ""
	f = nil
	f?(arg())
	must(trace == "", "statement nil skips")
	f = func(x int) *int { trace += "c"; return &one }
	f?(arg())
	must(trace == "ac", "statement discards result")
	trace = ""
	n?.Clear()
	n?.Multi(arg())
	must(trace == "", "skip void/multi statements")
	leaf?.Clear()
	leaf?.Multi(arg())
	must(trace == "zaq", "void/multi statements")
	trace = ""
	must(panics(func() { var missing map[int]*int; missing[key()] ??= fallback() }), "nil map store panic")
	must(trace == "kf", "nil map evaluates RHS before store")
	raceRead()
	fmt.Println("chains defaults dereferences assignments errors PASS")
}
func load(n *Node, b bool) (*int, error) { p := n?.Load(b)!; return p, nil }
func handle(n *Node, b bool) *int {
	return n?.Load(b) or err {
		trace += "h"
		return &one
	}
}
func numberDefault() int    { trace += "d"; return 9 }
func key() int              { trace += "k"; return 0 }
func getNode(n *Node) *Node { trace += "g"; return n }

func raceRead() {
	m := map[int]*int{0: &one}
	done := make(chan bool)
	go func() {
		for i := 0; i < 1000; i++ {
			if m[0] != &one {
				panic("read")
			}
		}
		done <- true
	}()
	for i := 0; i < 1000; i++ {
		m[0] ??= fallback()
	}
	<-done
}

func chooseKinds[T ~[]int | ~map[int]int](a, b T) T { return a ?? b }
func chooseAsAny[T ~*int](a, b T) any               { return a ?? b }
