// errorcheck

package p

type T struct {
	f float64
}

var t T

func F() {
	_ = complex(1.0) // ERROR "invalid operation|not enough arguments"
	_ = complex(t.f) // ERROR "invalid operation|not enough arguments"
}
