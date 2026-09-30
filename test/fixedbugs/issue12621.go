// run

// Issues 12576 and 12621: Negative untyped floating point constants
// with small magnitude round to 0, not negative zero.

package main

import "math"

var m = -1e-10000

func main() {
	if math.Signbit(m) {
		panic(m)
	}
}
