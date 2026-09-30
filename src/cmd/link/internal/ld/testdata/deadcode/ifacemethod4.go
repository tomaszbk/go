// Test that a live type's method is not live even if
// it matches an interface method, as long as the interface
// method is not used.

package main

type T int

//go:noinline
func (T) M() {}

type I interface{ M() }

var p *T
var pp *I

func main() {
	p = new(T)  // use type T
	pp = new(I) // use type I
	*pp = *p    // convert T to I, build itab
}
