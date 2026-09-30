// run cmplxdivide1.go


// Driver for complex division table defined in cmplxdivide1.go
// For details, see the comment at the top of cmplxdivide.c.

package main

import (
	"fmt"
	"math"
)

func calike(a, b complex128) bool {
	if imag(a) != imag(b) && !(math.IsNaN(imag(a)) && math.IsNaN(imag(b))) {
		return false
	}

	if real(a) != real(b) && !(math.IsNaN(real(a)) && math.IsNaN(real(b))) {
		return false
	}

	return true
}

func main() {
	bad := false
	for _, t := range tests {
		x := t.f / t.g
		if !calike(x, t.out) {
			if !bad {
				fmt.Printf("BUG\n")
				bad = true
			}
			fmt.Printf("%v/%v: got %v, want %v\n", t.f, t.g, x, t.out)
		}
	}
	if bad {
		panic("cmplxdivide failed.")
	}
}
