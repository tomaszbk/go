// compile

package main

type fun func()

func F[T any]() {
	_ = fun(func() {

	})
}
func main() {
	F[int]()
}
