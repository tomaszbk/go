// compile -d=ssa/check/on

package p

type S struct {
	i int
	a [1]struct{}
}

func foo(f func(S)) {
	defer f(func() S {
		return S{0, [1]struct{}{}}
	}())
}
