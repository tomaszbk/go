package a

type S struct{}

func callClosure(closure func()) {
	closure()
}

func (s *S) M() {
	callClosure(func() {
		defer f(s.m) // prevent closures to be inlined.
	})
}

func (s *S) m() {}

//go:noinline
func f(a ...any) {}
