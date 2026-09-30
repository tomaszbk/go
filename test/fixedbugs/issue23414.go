// compile

package p

var x struct{}

func f() bool {
	return x == x && x == x
}
