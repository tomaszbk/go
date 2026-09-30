// run

package main

func genfunc[T any](f func(c T)) {
	var r T

	f(r)
}

func myfunc(c string) {
	test2(c)
}

//go:noinline
func test2(a interface{}) {
	_ = a.(string)
}

func main() {
	genfunc(myfunc)
}
