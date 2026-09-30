package main

import "./a"

var m = a.I[int].M

var never bool

func main() {
	if never {
		m(nil)
	}
}
