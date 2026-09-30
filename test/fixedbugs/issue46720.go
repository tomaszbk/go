// compile


package p

func f() {
	nonce := make([]byte, 24)
	g((*[24]byte)(nonce))
}

//go:noinline
func g(*[24]byte) {}
