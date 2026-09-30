// errorcheck -0 -live


package p

type T struct { a, b, c, d string } // pass in registers, not SSA-able

//go:registerparams
func F() (r T) {
	r.a = g(1) // ERROR "live at call to g: r"
	r.b = g(2) // ERROR "live at call to g: r"
	r.c = g(3) // ERROR "live at call to g: r"
	r.d = g(4) // ERROR "live at call to g: r"
	return
}

func g(int) string
