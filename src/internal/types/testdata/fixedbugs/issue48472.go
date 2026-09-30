package p

func g() {
	var s string
	var i int
	_ = s /* ERROR "invalid operation: s + i (mismatched types string and int)" */ + i
}

func f(i int) int {
        i /* ERROR `invalid operation: i += "1" (mismatched types int and untyped string)` */ += "1"
        return i
}
