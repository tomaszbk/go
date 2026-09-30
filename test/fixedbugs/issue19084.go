// compile

// Issue 19084: SSA doesn't handle CONVNOP STRUCTLIT

package p

type T struct {
	a, b, c, d, e, f, g, h int // big, not SSA-able
}

func f() {
	_ = T(T{})
}
