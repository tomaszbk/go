// compile

// The gccgo compiler crashed while compiling a function that returned
// multiple zero-sized structs.
// https://gcc.gnu.org/PR80226.

package p

type S struct{}

func F() (S, S) {
	return S{}, S{}
}
