package bug

type Node interface {
	Eval(s *Scene)
}

type plug struct {
	node Node
}

type Scene struct {
	changed map[plug]bool
}
