// compile -d=libfuzzer


package p

func f(x, y int) {
	_ = x > y
	_ = y > x
}
