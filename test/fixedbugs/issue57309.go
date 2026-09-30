// run

package main

type I interface {
	M()
}

type S struct {
}

func (*S) M() {
}

func main() {
	func() {
		I(&S{}).M()
	}()
}
