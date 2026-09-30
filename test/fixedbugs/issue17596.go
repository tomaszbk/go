// compile


package foo

type T interface {
	foo()
}

func f() (T, int)

func g(v interface{}) (interface{}, int) {
	var x int
	v, x = f()
	return v, x
}
