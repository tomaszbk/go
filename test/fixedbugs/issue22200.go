// errorcheck


package p

func f1(x *[1<<30 - 1e6]byte) byte {
	sum := byte(0)
	for _, b := range *x {
		sum += b
	}
	return sum
}
func f2(x *[1<<30 + 1e6]byte) byte { // GC_ERROR "stack frame too large"
	sum := byte(0)
	for _, b := range *x {
		sum += b
	}
	return sum
}
