// run


package main

import "strings"

type S struct {
	a [3]int
}

//go:noinline
func f(i int) {
	var x [0]S
	x[i].a[1] = 3
}

func main() {
	defer func() {
		r := recover()
		if r == nil {
			panic("no panic (bug)")
		}
		got := r.(error).Error()
		if !strings.Contains(got, "index out of range [3] ") {
			panic("unexpected panic: " + got)
		}
	}()
	f(3)
}
