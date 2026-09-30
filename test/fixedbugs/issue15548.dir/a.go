package a

type I0 interface {
	I1
}

type T struct {
	I1
}

type I1 interface {
	M(*T) // removing * makes crash go away
}
