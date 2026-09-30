// compile


package p

func x() {
	func() func() {
		return func() {
			f := func() {}
			g, _ := f, 0
			g()
		}
	}()()
}
