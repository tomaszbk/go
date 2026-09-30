package nilfunc

func F() {}

func Comparison() {
	if F == nil { // ERROR "comparison of function F == nil is always false"
		panic("can't happen")
	}
}
