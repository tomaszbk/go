// compile

// https://gcc.gnu.org/PR101994
// gccgo compiler crash with zero-sized result.

package p

type Empty struct{}

func F() (int, Empty) {
	return 0, Empty{}
}
