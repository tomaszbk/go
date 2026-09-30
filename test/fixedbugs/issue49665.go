// run


package main

import "fmt"

var x any
var y interface{}

var _ = &x == &y // assert x and y have identical types

func main() {
	fmt.Printf("%T\n%T\n", &x, &y)
}
