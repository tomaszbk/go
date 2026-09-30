package a

func F() interface{} { return new(T[int]) }

type T[P any] int

func (x *T[P]) One() int { return x.Two() }
func (x *T[P]) Two() int { return 0 }
