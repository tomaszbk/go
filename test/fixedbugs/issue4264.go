// errorcheck


// issue 4264: reject int division by const 0

package main

func main() {
	var x int
	var y float64
	var z complex128

	println(x/0) // ERROR "division by zero"
	println(y/0)
	println(z/0)
}