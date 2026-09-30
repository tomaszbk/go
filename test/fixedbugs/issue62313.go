// compile


package p

func f() {
	var err error = nil
	defer func() { _ = &err }()
	err.Error()
}
