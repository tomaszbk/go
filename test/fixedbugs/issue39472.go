// compile -N


package p

func f(x float64) bool {
	x += 1
	return (x != 0) == (x != 0)
}
