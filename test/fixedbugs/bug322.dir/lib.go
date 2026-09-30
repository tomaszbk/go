package lib

type T struct {
	x int  // non-exported field
}

func (t T) M() {
}

func (t *T) PM() {
}
