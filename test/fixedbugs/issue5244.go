// run


// Issue 5244: the init order computation uses the wrong
// order for top-level blank identifier assignments.
// The example used to panic because it tries calling a
// nil function instead of assigning to f before.

package main

var f = func() int { return 1 }
var _ = f() + g()
var g = func() int { return 2 }

func main() {}
