package p

type T[U interface{ M() T /* ERROR "too many type arguments for type T" */ [U, int] }] int

type X int

func (X) M() T[X] { return 0 }
