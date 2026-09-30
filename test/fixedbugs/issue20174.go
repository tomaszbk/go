// compile -c=2

// Issue 20174: failure to typecheck contents of *T in the frontend.

package p

func f() {
	_ = (*interface{})(nil) // interface{} here used to not have its width calculated going into backend
	select {
	case _ = <-make(chan interface {
		M()
	}, 1):
	}
}
