package main

import (
	"./a"
	"fmt"
)

func BuildInt() int {
	return a.BuildInt()
}

func main() {
	if got, want := BuildInt(), 0; got != want {
		panic(fmt.Sprintf("got %d, want %d", got, want))
	}
}
