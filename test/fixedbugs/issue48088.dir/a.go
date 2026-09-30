package a

type T1 struct {
	*T2
}

type T2 struct {
}

func (t2 *T2) M() {
}

func F() {
	f(T1.M)
}

func f(f func(T1)) {
}
