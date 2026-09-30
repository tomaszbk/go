package cmplx

// Conj returns the complex conjugate of x.
func Conj(x complex128) complex128 { return complex(real(x), -imag(x)) }
