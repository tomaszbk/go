// compile

// Issue 52870: gofrontend gave incorrect error when incorrectly
// compiling ambiguous promoted method.

package p

type S1 struct {
	*S2
}

type S2 struct {
	T3
	T4
}

type T3 int32

func (T3) M() {}

type T4 int32

func (T4) M() {}
