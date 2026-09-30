// run


// Test static interface conversion of interface value nil.

package main

type R interface { R() }
type RW interface { R(); W() }

var e interface {}
var r R
var rw RW

func main() {
	r = r
	r = rw
	e = r
	e = rw
	rw = rw
}
