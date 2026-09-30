// compile

// gofrontend incorrectly gave an error for this code.

package p

type B bool

func main() {
	var v B = false
	if (true && true) && v {
	}
}
