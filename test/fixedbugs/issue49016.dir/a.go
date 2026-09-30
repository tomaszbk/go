package a

type Node interface {
	Position()
}

type noder struct{}

func (noder) Position() {}

type Scope map[int][]Node

func (s Scope) M1() Scope {
	if x, ok := s[0]; ok {
		return x[0].(struct {
			noder
			Scope
		}).Scope
	}
	return nil
}

func (s Scope) M2() Scope {
	if x, ok := s[0]; ok {
		st, _ := x[0].(struct {
			noder
			Scope
		})
		return st.Scope
	}
	return nil
}
