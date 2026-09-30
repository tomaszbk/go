package iface_i

type I interface {
	M()
}

type T struct {
}

func (t *T) M() {
}

// *T implements I
