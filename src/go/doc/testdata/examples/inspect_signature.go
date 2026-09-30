package foo_test

import (
	"bytes"
	"io"
)

func getReader() io.Reader { return nil }

func do(b bytes.Reader) {}

func Example() {
	getReader()
	do()
	// Output:
}

func ExampleIgnored() {
}
