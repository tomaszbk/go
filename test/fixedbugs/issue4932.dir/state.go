package state

import "./foo"

func Public() {
	var s Settings
	s.op()
}

type State struct{}

func (s *State) x(*Settings) {}

type Settings struct{}

func (c *Settings) x() {
	run([]foo.Op{{}})
}

func run([]foo.Op) {}

func (s *Settings) op() foo.Op {
	return foo.Op{}
}
