// compile

package main

var a, b, c interface{} = func() (_, _, _ int) { return 1, 2, 3 }()

func main() {
	println(a.(int), b.(int), c.(int))
}
