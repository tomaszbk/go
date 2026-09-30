package b

type t int

func (t) m() {}

func F1() interface{} { return struct{ t }{} }
func F2() interface{} { return *new(struct{ t }) }
func F3() interface{} { var x [1]struct{ t }; return x[0] }
