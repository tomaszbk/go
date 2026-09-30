// run


package main

type TestSuite struct {
	Tests []int
}

var Suites = []TestSuite{
	Dicts,
}
var Dicts = TestSuite{
	Tests: []int{0},
}

func main() {
	if &Dicts.Tests[0] != &Suites[0].Tests[0] {
		panic("bad")
	}
}
