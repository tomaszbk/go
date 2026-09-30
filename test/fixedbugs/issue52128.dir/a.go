package a

type I interface{}

type F func()

type s struct {
	f F
}

func NewWithF(f F) *s {
	return &s{f: f}
}

func NewWithFuncI(func() I) *s {
	return &s{}
}
