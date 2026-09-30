// compile

// https://gcc.gnu.org/PR89321
// gccgo compiler crash building map literals with a zero-sized value type.

package p

type M map[byte]struct{}

var (
	M1 = M{1: {}, 2: {}, 3: {}}
	M2 = M{1: {}, 2: {}}
)
