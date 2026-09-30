// compile

package p

func F() {
A:
	for range (func(func() bool))(nil) {
		_ = func() {
		A:
			goto A
		}
		goto A
	}
}
