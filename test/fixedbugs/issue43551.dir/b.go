package b

import "./a"

type S a.S
type Key a.Key

func (s S) A() Key {
	return Key(a.S(s).A())
}
