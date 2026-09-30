// compile

package p

func g[T any]() {
	type U []T
	type V []int
}

type S[T any] struct {
}

func (s S[T]) m() {
	type U []T
	type V []int
}

func f() {
	type U []int
}

type X struct {
}

func (x X) m() {
	type U []int
}
