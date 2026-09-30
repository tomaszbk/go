// compile

// Issue 77303: compiler crash on array of zero-size ASPECIAL elements.

package p

type zeroSizeSpecial struct {
	_ [0]float64
}

var x [3]zeroSizeSpecial

func f() bool {
	return x == x
}
