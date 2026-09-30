package a

const N = 2+3i

func Func() []complex128 {
	return []complex128{1, complex(2, 3), complex(4, 5)}
}

func Mul(z complex128) complex128 {
	return z * (3 + 4i)
}
