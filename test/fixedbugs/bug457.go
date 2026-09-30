// run


// Issue 4197: growing a slice of zero-width elements
// panics on a division by zero.

package main

func main() {
	var x []struct{}
	x = append(x, struct{}{})
}
