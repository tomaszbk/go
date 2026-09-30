// run


package main

func f[T any](i interface{}) {
	switch i.(type) {
	case T:
		println("T")
	case int:
		println("int")
	default:
		println("other")
	}
}

type myint int
func (myint) foo() {
}

func main() {
	f[interface{}](nil)
	f[interface{}](6)
	f[interface{foo()}](nil)
	f[interface{foo()}](7)
	f[interface{foo()}](myint(8))
}
