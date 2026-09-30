package main

import (
	"./a"
	"fmt"
)

func main() {
	const want = 120

	if got := a.Fact(5); got != want {
		panic(fmt.Sprintf("got %d, want %d", got, want))
	}

	if got := a.Fact[int64](5); got != want {
		panic(fmt.Sprintf("got %d, want %d", got, want))
	}

	if got := a.Fact(5.0); got != want {
		panic(fmt.Sprintf("got %f, want %f", got, want))
	}
}
