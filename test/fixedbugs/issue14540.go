// errorcheck


package p

func f(x int) {
	switch x {
	case 0:
		fallthrough
		; // ok
	case 1:
		fallthrough // ERROR "fallthrough statement out of place"
		{}
	case 2:
		fallthrough // ERROR "cannot fallthrough"
	}
}
