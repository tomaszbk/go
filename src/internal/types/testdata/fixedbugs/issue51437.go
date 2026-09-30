package p

type T struct{}

func (T) m() []int { return nil }

func f(x T) {
	for _, x := range func() []int {
		return x.m() // x declared in parameter list of f
	}() {
		_ = x // x declared by range clause
	}
}
