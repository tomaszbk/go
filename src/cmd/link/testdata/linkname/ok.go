// Use of public API is ok.

package main

import (
	"iter"
	"unique"
)

func seq(yield func(int) bool) {
	yield(123)
}

var s = "hello"

func main() {
	h := unique.Make(s)
	next, stop := iter.Pull(seq)
	defer stop()
	println(h.Value())
	println(next())
	println(next())
}
