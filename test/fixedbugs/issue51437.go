// compile


package p

type T struct{}

func (T) m() []T { return nil }

func f(x T) {
	for _, x := range func() []T {
		return x.m()
	}() {
		_ = x
	}
}
