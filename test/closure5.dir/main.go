// Check correctness of various closure corner cases
// that are expected to be inlined
package main

import "./a"

func main() {
	if !a.G()()() {
		panic("FAIL")
	}
}
