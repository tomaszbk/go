package main

import (
	"fmt"

	"./a"
)

var v = a.S{}

func main() {
	want := "{{ 0}}"
	if got := fmt.Sprint(v.F); got != want {
		panic(got)
	}
}
