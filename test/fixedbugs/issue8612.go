//compile


// Gccgo had a bug comparing a struct or array value with an interface
// values, when the struct or array was not addressable.

package p

type A [10]int

type S struct {
	i int
}

func F1() S {
	return S{0}
}

func F2() A {
	return A{}
}

func Cmp(v interface{}) bool {
	if F1() == v {
		return true
	}
	if F2() == v {
		return true
	}
	return false
}
