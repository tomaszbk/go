// errorcheck

package p

type T struct {
	f [1]int
}

func a() {
	_ = T // ERROR "type T is not an expression|invalid use of type|not an expression"
}

func b() {
	var v [len(T{}.f)]int // ok
	_ = v
}
