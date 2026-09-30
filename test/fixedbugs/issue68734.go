// compile

// The gofrontend had a bug handling panic of an untyped constant expression.

package issue68734

func F1() {
	panic(1 + 2)
}

func F2() {
	panic("a" + "b")
}
