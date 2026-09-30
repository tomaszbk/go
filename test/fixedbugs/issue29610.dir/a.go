package a

type I interface {
	M(init bool)
}

var V I

func init() {
	V = nil
}
