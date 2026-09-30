package p

func _(x *int) {
	_ = 0 < x // ERROR "invalid operation"
	_ = x < 0 // ERROR "invalid operation"
}
