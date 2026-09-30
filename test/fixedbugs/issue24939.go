// compile

package main

type T interface {
	M(P)
}

type M interface {
	F() P
}

type P = interface {
	// The compiler cannot handle this case. Disabled for now.
	// See issue #25838.
	// I() M
}

func main() {}
