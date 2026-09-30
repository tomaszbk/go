// run -goexperiment fieldtrack


package main

func main() {
	var i interface{} = new(T)
	if _, ok := i.(interface{ Bad() }); ok {
		panic("FAIL")
	}
}

type T struct{ U }

type U struct{}

//go:nointerface
func (*U) Bad() {}
