package p

type Mer interface { M() }

func F[T Mer](p *T) {
	p.M /* ERROR "p.M undefined" */ ()
}

type MyMer int

func (MyMer) M() {}

func _() {
	F(new(MyMer))
	F[Mer](nil)
}
