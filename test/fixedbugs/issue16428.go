// errorcheck

package p

var (
	b = [...]byte("abc") // ERROR "outside of array literal|outside a composite literal"
	s = len(b)
)
