// errorcheck -d=panic


package main

import "bytes"

type _ struct{ bytes.nonexist } // ERROR "unexported|undefined"

type _ interface{ bytes.nonexist } // ERROR "unexported|undefined|expected signature or type name"

func main() {
	var _ bytes.Buffer
	var _ bytes.buffer // ERROR "unexported|undefined"
}
