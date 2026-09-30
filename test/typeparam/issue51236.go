// run

package main

type I interface {
	[]byte
}

func F[T I]() {
	var t T
	explodes(t)
}

func explodes(b []byte) {}

func main() {

}
