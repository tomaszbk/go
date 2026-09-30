package a

type m struct {
	S string
}

var g = struct {
	m
	P string
}{
	m{"a"},
	"",
}

type S struct{}

func (s *S) M(p string) {
	r := g
	r.P = p
}
