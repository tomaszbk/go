// compile


package p

type Symbol interface{}

type Value interface {
	String() string
}

type Object interface {
	String() string
}

type Scope struct {
	outer *Scope
	elems map[string]Object
}

func (s *Scope) findouter(name string) (*Scope, Object) {
	return s.outer.findouter(name)
}

func (s *Scope) Resolve(name string) (sym Symbol) {
	if _, obj := s.findouter(name); obj != nil {
		sym = obj.(Symbol)
	}
	return
}

type ScopeName struct {
	scope *Scope
}

func (n *ScopeName) Get(name string) (Value, error) {
	return n.scope.Resolve(name).(Value), nil
}
