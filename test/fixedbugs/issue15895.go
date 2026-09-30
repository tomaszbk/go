// compile

// func bad used to fail to compile.

package p

type A [1]int

func bad(x A) {
	switch x {
	case A([1]int{1}):
	case A([1]int{1}):
	}
}

func good(x A) {
	y := A([1]int{1})
	z := A([1]int{1})
	switch x {
	case y:
	case z:
	}
}
