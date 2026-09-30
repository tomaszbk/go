// compile

// Issue 20145: some func types weren't dowidth-ed by the front end,
// leading to races in the backend.

package p

func f() {
	_ = (func())(nil)
}
