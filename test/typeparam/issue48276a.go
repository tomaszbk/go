// run


package main

import "fmt"

func main() {
	IsZero[int](0)
}

func IsZero[T comparable](val T) bool {
	var zero T
	fmt.Printf("%v:%v\n", zero, val)
	return val != zero
}
