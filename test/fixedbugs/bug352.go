// run

package main

var x [10][0]byte
var y = make([]struct{}, 10)

func main() {
	if &x[1] != &x[2] {
		println("BUG: bug352 [0]byte")
	}
	if &y[1] != &y[2] {
		println("BUG: bug352 struct{}")
	}
}
