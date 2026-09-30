// run

// issue 6269: name collision on method names for function local types.

package main

type foo struct{}

func (foo) Error() string {
	return "ok"
}

type bar struct{}

func (bar) Error() string {
	return "fail"
}

func unused() {
	type collision struct {
		bar
	}
	_ = collision{}
}

func main() {
	type collision struct {
		foo
	}
	s := error(collision{})
	if str := s.Error(); str != "ok" {
		println("s.Error() ==", str)
		panic(`s.Error() != "ok"`)
	}
}
