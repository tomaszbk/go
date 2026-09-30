// compile


// Issue 13262: cmd/compile: bogus "fallthrough
// statement out of place" error

package p

func f() int {
	var a int
	switch a {
	case 0:
		return func() int { return 1 }()
		fallthrough
	default:
	}
	return 0
}
