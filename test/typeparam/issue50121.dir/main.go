package main

import (
	"./a"
)

//go:noinline
func BuildInt() int {
	return a.BuildInt()
}

func main() {
	BuildInt()
}
