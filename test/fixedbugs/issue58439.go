// compile


package p

var x = f(-1)
var y = f(64)

func f(x int) int {
	return 1 << x
}
