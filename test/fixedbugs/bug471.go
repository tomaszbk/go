// compile

// Caused an internal compiler error in gccgo.

package p

type C chan struct{}

func (c C) F() {
	select {
	case c <- struct{}{}:
	default:
	}
}
