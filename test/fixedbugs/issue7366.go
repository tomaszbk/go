// compile

// issue 7366: generates a temporary with ideal type
// during comparison of small structs.

package main

type T struct {
	data [10]byte
}

func main() {
	var a T
	var b T
	if a == b {
	}
}
