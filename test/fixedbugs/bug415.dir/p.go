package p

type A struct {
	s struct{int}
}

func (a *A) f() {
	a.s = struct{int}{0}
}

