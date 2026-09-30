// compile


// PR61248: Transformations to recover calls made them fail typechecking in gccgo.

package main

func main() {
	var f func(int, interface{})
	go f(0, recover())
}
