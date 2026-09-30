// run

package main

var p *int

func main() {
	var i int
	p = &i // escape i to keep the compiler from making the closure trivial

	func() { i++ }()
}
