// compile


package p

func f() {
	if false {
		defer func() {
			_ = recover()
		}()
	}
}

func g() {
	for false {
		defer func() {
			_ = recover()
		}()
	}
}
