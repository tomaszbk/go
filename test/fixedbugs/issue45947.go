// compile


package p

func f() {
	_ = func() func() {
		return func() {
		l:
			goto l
		}
	}()
}
