package math

// Copysign returns a value with the magnitude of f
// and the sign of sign.
func Copysign(f, sign float64) float64 {
	return Float64frombits(Float64bits(f)&^signMask | Float64bits(sign)&signMask)
}
