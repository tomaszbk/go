package p

type A struct{}

func (*A) m() int { return 0 }

var _ = A.m /* ERROR "invalid method expression A.m (needs pointer receiver (*A).m)" */ ()
var _ = (*A).m(nil)

type B struct{ A }

var _ = B.m // ERROR "invalid method expression B.m (needs pointer receiver (*B).m)"
var _ = (*B).m

var _ = struct{ A }.m // ERROR "invalid method expression struct{A}.m (needs pointer receiver (*struct{A}).m)"
