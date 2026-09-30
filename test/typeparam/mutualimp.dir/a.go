package a

type X int

func (x X) M() X { return x }

func F[T interface{ M() U }, U interface{ M() T }]() {}
func G()                                             { F[X, X]() }
