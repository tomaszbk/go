// errorcheck


package p

type t struct {
	x int // GCCGO_ERROR "duplicate field name .x."
	x int // GC_ERROR "duplicate field x|x redeclared"
}

func f(t *t) int {
	return t.x
}
