// compile

package foo

type T struct {
	x int
	_ int
}

func main() {
	_ = T{0, 0}

	x := T{1, 1}
	_ = x
}
