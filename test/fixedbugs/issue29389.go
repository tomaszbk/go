// compile

// Make sure we can correctly compile method expressions
// where the method is implicitly declared.

package main

import "io"

func main() {
	err := io.EOF
	_ = err.Error
}
