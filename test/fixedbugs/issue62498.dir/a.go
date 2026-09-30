package a

func One(L any) {
	func() {
		defer F(L)
	}()
}

func F(any) {}
