package a

type Builder struct {
	x int
}

func (tb Builder) Build() (out struct {
	x interface{}
	s string
}) {
	out.x = nil
	out.s = "hello!"
	return
}
