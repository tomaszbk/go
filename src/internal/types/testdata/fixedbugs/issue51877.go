package p

type S struct {
	f1 int
	f2 bool
}

var (
	_ = S{0}                    /* ERROR "too few values in struct literal" */
	_ = struct{ f1, f2 int }{0} /* ERROR "too few values in struct literal" */

	_ = S{0, true, "foo" /* ERROR "too many values in struct literal" */}
	_ = struct{ f1, f2 int }{0, 1, 2 /* ERROR "too many values in struct literal" */}
)
