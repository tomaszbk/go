// compile


// Used to crash when compiling functions containing
// forward refs in dead code.

package p

var f func(int)

func g() {
l1:
	i := 0
	goto l1
l2:
	f(i)
	goto l2
}
