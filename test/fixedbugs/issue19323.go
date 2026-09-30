// errorcheck


package p

func g() {}

func f() {
	g()[:] // ERROR "g.* used as value|attempt to slice object that is not"
}

func g2() ([]byte, []byte) { return nil, nil }

func f2() {
	g2()[:] // ERROR "multiple-value g2.* in single-value context|attempt to slice object that is not|2\-valued g"
}
