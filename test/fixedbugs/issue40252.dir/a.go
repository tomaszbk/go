package a

type I interface {
	Func()
}

func Call() {
	f := I.Func
	f(nil)
}
