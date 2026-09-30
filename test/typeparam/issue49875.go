// compile


package p

func f(args ...interface{}) {}

func g() {
	var args []any
	f(args...)
}
