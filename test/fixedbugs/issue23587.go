// errorcheck


package p

func _(x int) {
	_ = ~x    // unary ~ permitted but the type-checker will complain
}

func _(x int) {
	_ = x ~ x // ERROR "unexpected ~ at end of statement"
}
