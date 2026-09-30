package p

func f(s string) int {
	for range s {
	}
} // ERROR "missing return"
