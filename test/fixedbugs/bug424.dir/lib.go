package lib

type I interface {
	m() string
}

type T struct{}

// m is not accessible from outside this package.
func (t *T) m() string {
	return "lib.T.m"
}
