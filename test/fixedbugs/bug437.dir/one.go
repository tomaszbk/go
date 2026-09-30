package one

type I1 interface {
	f()
}

type S1 struct {
}

func (s S1) f() {
}

func F1(i1 I1) {
}
