// compile


package p

func f(b bool, c complex128) func(complex128) complex128 {
	return func(p complex128) complex128 {
		b = (p+1i == 0) && b
		return (p + 2i) * (p + 3i - c)
	}
}
