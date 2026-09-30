// compile -c=4


// Issue 20162: embedded interfaces weren't dowidth-ed by the front end,
// leading to races in the backend.

package p

func Foo() {
	_ = (make([]func() interface {
		M(interface{})
	}, 1))
}
