// errorcheck


// Expects to see error messages on 'p' exponents
// for non-hexadecimal floats.

package main

import "fmt"

const (
	x1 = 1.1    // float
	x2 = 1e10   // float
	x3 = 0x1e10 // integer (e is a hex digit)
)

const x4 = 0x1p10 // valid hexadecimal float
const x5 = 1p10   // ERROR "'p' exponent requires hexadecimal mantissa|invalid prefix"
const x6 = 0P0    // ERROR "'P' exponent requires hexadecimal mantissa|invalid prefix"

func main() {
	fmt.Printf("%g %T\n", x1, x1)
	fmt.Printf("%g %T\n", x2, x2)
	fmt.Printf("%g %T\n", x3, x3)
	fmt.Printf("%g %T\n", x4, x4)
	fmt.Printf("%g %T\n", x5, x5)
	fmt.Printf("%g %T\n", x6, x6)
}
