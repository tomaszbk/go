package p

func _(x any) {
	switch x {
	case 0:
		fallthrough // ERROR "fallthrough statement out of place"
		_ = x
	default:
	}

	switch x.(type) {
	case int:
		fallthrough // ERROR "cannot fallthrough in type switch"
	default:
	}
}
