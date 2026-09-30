// run


package main

import . "testing"  // defines file-level T

type _ B // make use of package "testing" (but don't refer to T)

type S struct {
	T int
}

func main() {
	_ = &S{T: 1}	// should work
}
