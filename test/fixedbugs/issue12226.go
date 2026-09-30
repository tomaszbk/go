// run


package main

import "fmt"

func main() {
	if []byte("foo")[0] == []byte("b")[0] {
		fmt.Println("BUG: \"foo\" and \"b\" appear to have the same first byte")
	}
}
