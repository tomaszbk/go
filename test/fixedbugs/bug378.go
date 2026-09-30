// run


// Issue 2497

package main

type Header struct{}
func (h Header) Method() {}

var _ interface{} = Header{}

func main() {
  	type X Header
  	var _ interface{} = X{}
}
