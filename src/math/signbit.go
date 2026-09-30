package math

// Signbit reports whether x is negative or negative zero.
func Signbit(x float64) bool {
	return int64(Float64bits(x)) < 0
}
