package main

import "./a"

func F() {
	// store 1 in a.global
	x, y := 1, 2
	t := a.T{Pointer: &x}
	a.Store(&t)
	_ = y
}

func G() {
	// store 4 in a.global2
	x, y := 3, 4
	t := a.T{Pointer: &y}
	a.Store2(&t)
	_ = x
}

func main() {
	F()
	G()
	p := a.Get()
	n := *p
	if n != 1 {
		println(n, "!= 1")
		panic("n != 1")
	}
}
