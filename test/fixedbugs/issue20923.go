// compile

// Issue 20923: gccgo failed to compile parenthesized select case expressions.

package p

func F(c chan bool) {
	select {
	case (<-c):
	case _ = (<-c):
	case _, _ = (<-c):
	case (c) <- true:
	default:
	}
}
