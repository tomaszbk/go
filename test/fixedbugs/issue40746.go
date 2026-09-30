// compile

package p

func f(x byte, b bool) byte {
	var c byte
	if b {
		c = 1
	}

	if int8(c) < 0 {
		x++
	}
	return x
}
