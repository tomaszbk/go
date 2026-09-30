package cmplx

import "math"

// Rect returns the complex number x with polar coordinates r, θ.
func Rect(r, θ float64) complex128 {
	s, c := math.Sincos(θ)
	return complex(r*c, r*s)
}
