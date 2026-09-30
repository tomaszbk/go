// compile

// golang.org/issue/6298.
// Used to cause "internal error: typename ideal bool"

package main

func main() {
	var x interface{} = "abc"[0] == 'a'
	_ = x
}
