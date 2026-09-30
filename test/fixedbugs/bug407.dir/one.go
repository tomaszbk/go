package one

// Issue 2877
type T struct {
	f func(t *T, arg int)
	g func(t T, arg int)
}

func (t *T) foo(arg int) {}
func (t T) goo(arg int) {}

func (t *T) F() { t.f = (*T).foo }
func (t *T) G() { t.g = T.goo }



