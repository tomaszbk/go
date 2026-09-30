// run

package main


func f[T comparable](i any) {
	var t T

	if i != t {
		println("FAIL: if i != t")
	}
}

type myint int

func (m myint) foo() {
}

type fooer interface {
	foo()
}

type comparableFoo interface {
	comparable
	foo()
}

func g[T comparableFoo](i fooer) {
	var t T

	if i != t {
		println("FAIL: if i != t")
	}
}

func main() {
	f[int](int(0))
	g[myint](myint(0))
}
