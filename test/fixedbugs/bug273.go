// run

// https://golang.org/issue/589

package main

import "unsafe"

var bug = false

var minus1 = -1
var five = 5
var big int64 = 10 | 1<<46

type block [1 << 19]byte

var g1 []block

func shouldfail(f func(), desc string) {
	defer func() { recover() }()
	f()
	if !bug {
		println("BUG")
		bug = true
	}
	println("didn't crash: ", desc)
}

func badlen() {
	g1 = make([]block, minus1)
}

func biglen() {
	g1 = make([]block, big)
}

func badcap() {
	g1 = make([]block, 10, minus1)
}

func badcap1() {
	g1 = make([]block, 10, five)
}

func bigcap() {
	g1 = make([]block, 10, big)
}

type cblock [1<<16 - 1]byte

var g4 chan cblock

func badchancap() {
	g4 = make(chan cblock, minus1)
}

func bigchancap() {
	g4 = make(chan cblock, big)
}

func overflowchan() {
	const ptrSize = unsafe.Sizeof(uintptr(0))
	g4 = make(chan cblock, 1<<(30*(ptrSize/4)))
}

func main() {
	shouldfail(badlen, "badlen")
	shouldfail(biglen, "biglen")
	shouldfail(badcap, "badcap")
	shouldfail(badcap1, "badcap1")
	shouldfail(bigcap, "bigcap")
	shouldfail(badchancap, "badchancap")
	shouldfail(bigchancap, "bigchancap")
	shouldfail(overflowchan, "overflowchan")
}
