package a

type T struct{ _ int }
func (t T) M() {}

type I interface { M() }

func F() {
	var t I = &T{}
	t.M() // call to the wrapper (*T).M
}
