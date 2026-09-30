// compile -B


// Make sure that we can at least compile this code
// successfully with -B. We can't ever produce the right
// answer at runtime with -B, as the access must panic.

package p

type A [0]byte

func (a *A) Get(i int) byte {
	return a[i]
}
