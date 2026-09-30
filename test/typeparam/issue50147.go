// compile


package p

func Foo[T any, U interface{ *T }](x T) {
	var _ U = &x
}
