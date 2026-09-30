package p

type A interface {
	a()
}

type AB interface {
	A
	b()
}

type AAB struct {
	A
	AB
}

var _ AB = AAB /* ERROR "ambiguous selector AAB.a" */ {}
