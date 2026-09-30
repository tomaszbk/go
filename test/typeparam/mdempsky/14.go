// run


package main

// Zero returns the zero value of T
func Zero[T any]() (_ T) {
	return
}

type AnyInt[X any] int

func (AnyInt[X]) M() {
	var have interface{} = Zero[X]()
	var want interface{} = Zero[MyInt]()

	if have != want {
		println("FAIL")
	}
}

type I interface{ M() }

type MyInt int
type U = AnyInt[MyInt]

var x = U(0)
var i I = x

func main() {
	x.M()
	U.M(x)
	(*U).M(&x)

	i.M()
	I.M(x)
}
