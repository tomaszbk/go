// compile

// Used to crash when compiling assignments involving [0]T,
// where T is not SSA-able.

package p

type s struct {
	a, b, c, d, e int
}

func f() {
	var i int
	arr := [0]s{}
	arr[i].a++
}
