package cmplx

// Polar returns the absolute value r and phase θ of x,
// such that x = r * e**θi.
// The phase is in the range [-Pi, Pi].
func Polar(x complex128) (r, θ float64) {
	return Abs(x), Phase(x)
}
