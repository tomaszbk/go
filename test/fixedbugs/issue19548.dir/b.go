package main

import "./a"

type Value interface {
	a.Stringer
	Addr() *a.Mode
}

var global a.Mode

func f() int {
	var v Value
	v = &global
	return int(v.String()[0])
}

func main() {
	f()
}
