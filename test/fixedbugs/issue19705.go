// compile


package p

func f1() {
	f2()
}

func f2() {
	if false {
		_ = func() {}
	}
}
