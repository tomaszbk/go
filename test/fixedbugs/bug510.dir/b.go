package main

import "./a"

func main() {
	_, ok := a.F().(*map[int]bool)
	if !ok {
		panic("bad type")
	}
}
