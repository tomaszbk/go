// compile


// Issue 10654: Failure to use generated temps
// for function calls etc. in boolean codegen.

package main

var s string

func main() {
	if (s == "this") != (s == "that") {
	}
}
