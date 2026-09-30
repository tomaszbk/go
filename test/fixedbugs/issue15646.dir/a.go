package a

type T struct{}

func (T) m() string {
	return "m"
}

func (*T) mp() string {
	return "mp"
}

func F() func(T) string {
	return T.m // method expression
}

func Fp() func(*T) string {
	return (*T).mp // method expression
}
