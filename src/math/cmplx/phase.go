package cmplx

import "math"

// Phase returns the phase (also called the argument) of x.
// The returned value is in the range [-Pi, Pi].
func Phase(x complex128) float64 { return math.Atan2(imag(x), real(x)) }
