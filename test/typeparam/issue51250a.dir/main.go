package main

import (
	"./a"
	"./b"
)

func main() {
	switch b.I.(type) {
	case a.G[b.T]:
	case int:
		panic("bad")
	case float64:
		panic("bad")
	default:
		panic("bad")
	}

	b.F(a.G[b.T]{})
}
