package p

type A = uint8
type E uint8

func f[P ~A](P) {}

func g(e E) {
	f(e)
}
