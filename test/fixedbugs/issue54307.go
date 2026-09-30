// compile


package p

func f[Int int, Uint uint]() {
	_ = uint(Int(-1))
	_ = uint(Uint(0) - 1)
}

func g[String string]() {
	_ = String("")[100]
}

var _ = f[int, uint]
var _ = g[string]
