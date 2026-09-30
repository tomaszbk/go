// compile

// The compiler used the name "glob" as the function holding a global
// function literal, colliding with an actual function named "glob".

package main

func glob() {
	func() {
	}()
}

var c1 = func() {
}

var c2 = func() {
}

func main() {
	glob()
	c1()
	c2()
}
