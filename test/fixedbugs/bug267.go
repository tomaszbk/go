// compile

package bug267

type T []int

var a []bool

func f1() {
	if a[T{42}[0]] {
	}
	// if (a[T{42}[0]]) {}  // this compiles
}

/*
6g bugs/bug267.go
bugs/bug267.go:14: syntax error: unexpected {, expecting :
*/
