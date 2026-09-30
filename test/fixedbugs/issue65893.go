// compile


package p

type (
	s  = struct{ f func(s1) }
	s1 = struct{ i I }
)

type I interface {
	S() *s
}
