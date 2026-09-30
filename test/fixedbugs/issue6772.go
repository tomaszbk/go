// errorcheck


package p

func f1() {
	for a, a := range []int{1, 2, 3} { // ERROR "a.* repeated on left side of :=|a redeclared"
		println(a)
	}
}

func f2() {
	var a int
	for a, a := range []int{1, 2, 3} { // ERROR "a.* repeated on left side of :=|a redeclared"
		println(a)
	}
	println(a)
}
