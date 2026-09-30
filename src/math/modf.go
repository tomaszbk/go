package math

// Modf returns integer and fractional floating-point numbers
// that sum to f. Both values have the same sign as f.
//
// Special cases are:
//
//	Modf(±Inf) = ±Inf, NaN
//	Modf(NaN) = NaN, NaN
func Modf(f float64) (integer float64, fractional float64) {
	integer = Trunc(f)
	fractional = Copysign(f-integer, f)
	return
}
