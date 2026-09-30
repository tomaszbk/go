package a

func Start() interface{ Stop() } {
	return new(Stopper)
}

type Stopper struct{}

func (s *Stopper) Stop() {}
