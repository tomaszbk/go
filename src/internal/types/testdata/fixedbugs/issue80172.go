package p

type List[P /* ERROR "instantiation cycle" */ any] struct{}

func (_ List[P]) m() (_ List[List[P]]) { return }
