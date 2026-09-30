// run -gcflags=all=-d=checkptr


package main

import (
	"regexp"
	"unique"
)

var dataFileRegexp = regexp.MustCompile(`^data\.\d+\.bin$`)

func main() {
	_ = dataFileRegexp
	unique.Make("")
}
