package main

import (
	"./b"
	"./c"
)

func main() {
	if b.G() != c.G() {
		println(b.G(), c.G())
		panic("bad")
	}
	if b.F() != c.F() {
		println(b.F(), c.F())
		panic("bad")
	}
}
