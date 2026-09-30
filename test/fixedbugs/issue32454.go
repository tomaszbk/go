// compile


package p

type T struct {
	s string
	f float64
}

func f() {
	var f float64
	var st T
	for {
		switch &st.f {
		case &f:
			f = 1
		}
	}
}
