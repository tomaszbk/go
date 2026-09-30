package main

import (
	"os"
	"prog/dep"
)

//go:noinline
func fifth() {
	println("hubba")
}

//go:noinline
func sixth() {
	println("wha?")
}

func main() {
	println(dep.Dep1())
	if len(os.Args) > 1 {
		fifth()
	} else {
		sixth()
	}
}
