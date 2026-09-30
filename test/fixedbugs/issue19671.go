// compile

// Used to crash when compiling assignments involving [0]T,
// where T is not SSA-able.

package a

func f() {
	var i int
	arr := [0][2]int{}
	arr[i][0] = 0
}
