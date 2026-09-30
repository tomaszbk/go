package issue25596

type E interface {
	M() T
}

type T interface {
	E
}
