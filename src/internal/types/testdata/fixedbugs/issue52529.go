package p

type Foo[P any] struct {
	_ *Bar[P]
}

type Bar[Q any] Foo[Q]

func (v *Bar[R]) M() {
	_ = (*Foo[R])(v)
}
