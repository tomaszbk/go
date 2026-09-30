package p1

type Magic int

type T struct {
	x interface{}
}

func (t *T) M() bool {
	_, ok := t.x.(Magic)
	return ok
}

func F(t *T) {
	println(t)
}
