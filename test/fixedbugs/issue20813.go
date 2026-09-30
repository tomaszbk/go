// errorcheck


package p

func f() {
	1 = 2 // ERROR "cannot assign to 1|invalid left hand side"
}
