package main

import (
	"fmt"
	"unsafe"
)

func choose[T ~*int](p, q T) T {
	if p == nil {
		return q
	}
	return p
}

var initial *int = func() *int {
	var p *int
	if p == nil {
		return &one
	}
	return p
}()

func scenarios() {
	leaf := &Node{P: &one, Count: 7, Items: []*int{&one}, Value: &one, Fn: func(x int) *int { return &one }}
	for _, n := range []*Node{nil, {Next: nil}, {Next: leaf}} {
		trace = ""
		var p *int
		t := root(n)
		if t != nil {
			t = t.Next
			if t != nil {
				p = t.P
			}
		}
		want := (*int)(nil)
		if n != nil && n.Next != nil {
			want = n.Next.P
		}
		must(p == want && trace == "r", "chain guards")
		count := 4
		if n != nil && n.Next != nil {
			count = n.Next.Count
		}
		wantCount := 4
		if n != nil && n.Next != nil {
			wantCount = n.Next.Count
		}
		must(count == wantCount, "scalar absence")
		trace = ""
		a := first()
		t = root(n)
		p = nil
		if t != nil {
			t = t.Next
			if t != nil {
				p = t.Get(arg())
			}
		}
		if p == nil {
			p = fallback()
		}
		p = collect(a, p, last())
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
	var p *int
	if f != nil {
		p = f(arg())
	}
	must(p == nil && trace == "", "nil safe call")
	f = leaf.Fn
	p = nil
	if f != nil {
		p = f(arg())
	}
	must(p == &one && trace == "a", "safe function call")
	var n *Node
	trace = ""
	p = nil
	if n != nil {
		p = n.NilMethod()
	}
	must(p == nil && trace == "", "skip nil receiver")
	must(n.NilMethod() == &two && trace == "n", "ordinary nil receiver")
	var i Reader = n
	trace = ""
	p = nil
	if i != nil {
		p = i.NilMethod()
	}
	must(p == &two && trace == "n", "typed nil interface")
	i = nil
	trace = ""
	p = nil
	if i != nil {
		p = i.NilMethod()
	}
	must(p == nil && trace == "", "nil interface")
	var method func() *int
	if n != nil {
		method = n.NilMethod
	}
	must(method == nil, "nil method value")
	method = nil
	if leaf != nil {
		method = leaf.NilMethod
	}
	must(method() == &one, "present method value")
	p = nil
	if n != nil {
		p = n.ByValue()
	}
	must(p == nil, "value receiver nil guard")
	must(panics(func() {
		t := &Node{}
		if t != nil {
			_ = t.Next.P
		}
	}), "unguarded tail panics")
	must(panics(func() {
		var t *Node
		if n != nil {
			t = n.Next
		}
		_ = t.P
	}), "parenthesized boundary")
	must(panics(func() {
		t := &Embedded{}
		if t != nil {
			_ = t.P
		}
	}), "promoted embedded pointer panics")
	p = nil
	if leaf != nil {
		p = leaf.Items[0]
	}
	must(p == &one, "guard index")
	p = nil
	if n != nil {
		p = n.Items[9]
	}
	must(p == nil, "skip index")
	p = nil
	if leaf != nil {
		p = leaf.Items[:1][0]
	}
	must(p == &one, "guard slice")
	p = nil
	if leaf != nil {
		p = leaf.Value.(*int)
	}
	must(p == &one, "guard assertion")
	p = nil
	if n != nil {
		p = n.Value.(*int)
	}
	must(p == nil, "skip assertion")
	var e *Errors
	var err error
	if e != nil {
		err = e.E
	}
	must(err == nil, "absent interface target")
	e = &Errors{}
	err = nil
	if e != nil {
		err = e.E
	}
	must(err != nil, "present typed nil interface target")
	var chosenErr error
	if e != nil && e.E != nil {
		chosenErr = e.E
	} else {
		chosenErr = sentinel
	}
	must(chosenErr == sentinel, "chain nil test before conversion")
	trace = ""
	var x int
	if leaf != nil {
		x = leaf.Count
	} else {
		x = numberDefault()
	}
	must(x == 7 && trace == "", "skip scalar fallback")
	if n != nil {
		x = n.Count
	} else {
		x = numberDefault()
	}
	must(x == 9 && trace == "d", "scalar fallback")
	var ptr *int
	must(deref(ptr, 10) == 10, "guarded deref")
	ptr = &one
	must(deref(ptr, 10) == 1, "present deref")
	var pp **int
	must(deref2(pp, 10) == 10, "nested deref")
	pp = &ptr
	must(deref2(pp, 10) == 1, "present nested deref")
	must(derefNode(n, 10) == 10, "deref chain")
	var pe *error
	must(derefError(pe) == sentinel, "deref nil interface")
	var emptyError error
	pe = &emptyError
	must(derefError(pe) == sentinel, "deref nil value")
	var s []int
	must(len(sliceDefault(s)) == 1, "slice nil")
	s = []int{}
	must(len(sliceDefault(s)) == 0, "slice empty")
	var m map[int]int
	must(len(mapDefault(m)) == 1, "map nil")
	m = map[int]int{}
	must(len(mapDefault(m)) == 0, "map empty")
	var ch chan int
	ch2 := make(chan int)
	must(channelDefault(ch, ch2) == ch2, "channel")
	var fn func() int
	if fn == nil {
		fn = func() int { return 4 }
	}
	must(fn() == 4, "function")
	var u unsafe.Pointer
	must(unsafeDefault(u, unsafe.Pointer(&one)) == unsafe.Pointer(&one), "unsafe pointer")
	var anyNil any = (*int)(nil)
	must(anyDefault(anyNil, any(&two)) == anyNil, "interface nil test")
	ptr = nil
	var boxed any
	if ptr != nil {
		boxed = ptr
	} else {
		boxed = &two
	}
	must(boxed == &two, "nil test before target conversion")
	trace = ""
	p = ptr
	if p == nil {
		p = fallback()
		if p == nil {
			p = &one
		}
	}
	must(p == &two && trace == "f", "right associative")
	must(choose((*int)(nil), &one) == &one, "generic")
	must(len(chooseKinds([]int(nil), []int{1})) == 1 && len(chooseKinds(map[int]int(nil), map[int]int{1: 1})) == 1, "mixed nilable type set")
	must(chooseAsAny((*int)(nil), &one) == &one, "generic target conversion")
	trace = ""
	slots := map[int]*int{}
	k := key()
	if slots[k] == nil {
		slots[k] = fallback()
	}
	must(slots[0] == &two && trace == "kf", "map coalescing assignment")
	trace = ""
	k = key()
	if slots[k] == nil {
		slots[k] = fallback()
	}
	must(trace == "k", "no redundant store")
	trace = ""
	g := getNode(leaf)
	if g.P == nil {
		g.P = fallback()
	}
	must(trace == "g", "field assignment operands")
	leaf.P = nil
	trace = ""
	g = getNode(leaf)
	if g.P == nil {
		g.P = fallback()
	}
	must(trace == "gf" && leaf.P == &two, "field lazy assignment")
	var local *int
	if local == nil {
		local = &one
	}
	if local == nil {
		local = &two
	}
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
	if f != nil {
		f(arg())
	}
	must(trace == "", "statement nil skips")
	f = func(x int) *int { trace += "c"; return &one }
	if f != nil {
		f(arg())
	}
	must(trace == "ac", "statement discards result")
	trace = ""
	if n != nil {
		n.Clear()
		n.Multi(arg())
	}
	must(trace == "", "skip void/multi statements")
	if leaf != nil {
		leaf.Clear()
		leaf.Multi(arg())
	}
	must(trace == "zaq", "void/multi statements")
	trace = ""
	must(panics(func() {
		var missing map[int]*int
		k := key()
		if missing[k] == nil {
			missing[k] = fallback()
		}
	}), "nil map store panic")
	must(trace == "kf", "nil map evaluates RHS before store")
	raceRead()
	fmt.Println("chains defaults dereferences assignments errors PASS")
}
func load(n *Node, b bool) (*int, error) {
	var p *int
	if n != nil {
		var err error
		p, err = n.Load(b)
		if err != nil {
			return nil, err
		}
	}
	return p, nil
}
func handle(n *Node, b bool) *int {
	var p *int
	if n != nil {
		var err error
		p, err = n.Load(b)
		if err != nil {
			trace += "h"
			return &one
		}
	}
	return p
}
func numberDefault() int    { trace += "d"; return 9 }
func key() int              { trace += "k"; return 0 }
func getNode(n *Node) *Node { trace += "g"; return n }

func deref(p *int, d int) int {
	if p == nil {
		return d
	}
	return *p
}
func deref2(p **int, d int) int {
	if p == nil || *p == nil {
		return d
	}
	return **p
}
func derefNode(p *Node, d int) int {
	if p == nil || p.P == nil {
		return d
	}
	return *p.P
}
func derefError(p *error) error {
	if p == nil || *p == nil {
		return sentinel
	}
	return *p
}
func sliceDefault(s []int) []int {
	if s == nil {
		return []int{1}
	}
	return s
}
func mapDefault(m map[int]int) map[int]int {
	if m == nil {
		return map[int]int{1: 1}
	}
	return m
}
func channelDefault(a, b chan int) chan int {
	if a == nil {
		return b
	}
	return a
}
func unsafeDefault(a, b unsafe.Pointer) unsafe.Pointer {
	if a == nil {
		return b
	}
	return a
}
func anyDefault(a, b any) any {
	if a == nil {
		return b
	}
	return a
}

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
		if m[0] == nil {
			m[0] = fallback()
		}
	}
	<-done
}

func chooseKinds[T ~[]int | ~map[int]int](a, b T) T {
	if a != nil {
		return a
	}
	return b
}
func chooseAsAny[T ~*int](a, b T) any {
	if a != nil {
		return a
	}
	return b
}
