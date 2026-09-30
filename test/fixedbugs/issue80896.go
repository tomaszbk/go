// errorcheck -std

//go:build !(386 || arm || mips || mipsle)


package main

const N = 2e6

type Big = [4 * 3 * N * N]int

func sink(x Big) {} // ERROR "stack frame too large"

func h(x0, x1, x2 Big) { // ERROR "stack frame too large"
	sink(x0)
	sink(x1)
	sink(x2)
}
