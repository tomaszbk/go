package main

import (
	"./a"
	"./b"
)

func main() {
	if a.F() == b.F() {
		panic("FAIL")
	}
	if a.X == b.X {
		panic("FAIL")
	}
}
