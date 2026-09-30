package main

import "./a"

var g = a.G()

func main() {
	if !a.F() {
		panic("FAIL")
	}
	if !g() {
		panic("FAIL")
	}
}
