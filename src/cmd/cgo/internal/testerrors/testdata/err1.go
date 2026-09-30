package main

/*
#cgo LDFLAGS: -L/nonexist

void test() {
	xxx;		// ERROR HERE
}

// Issue 8442.  Cgo output unhelpful error messages for
// invalid C preambles.
void issue8442foo(UNDEF*); // ERROR HERE
*/
import "C"

func main() {
	C.test()
}
