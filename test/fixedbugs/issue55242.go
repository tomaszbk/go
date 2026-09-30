// compile

// Issue 55242: gofrontend crash calling function that returns
// trailing empty struct.

package p

func F1() (int, struct{}) {
	return 0, struct{}{}
}

func F2() {
	F1()
}
