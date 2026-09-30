// compile


// issue 7346 : internal error "doasm" error due to checknil
// of a nil literal.

package main

func main() {
	_ = *(*int)(nil)
}
