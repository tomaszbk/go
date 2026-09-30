// compile

// /tmp/x.go:5: illegal constant expression: bool == interface {}

package p

var m = map[interface{}]struct{}{
	nil:  {},
	true: {},
}
