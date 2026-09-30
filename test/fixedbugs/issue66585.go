// run


package main

var x = 0
var a = foo()
var b = x

func foo() int {
	x++
	return x
}

func main() {
	if a != 1 {
		panic("unexpected a value")
	}
	if b != 1 {
		panic("unexpected b value")
	}
}
