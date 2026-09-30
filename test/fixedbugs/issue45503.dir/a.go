package a

type S struct{}

func (s *S) M() {
	s.m((*S).N)
}

func (s *S) N() {}

func (s *S) m(func(*S)) {}
