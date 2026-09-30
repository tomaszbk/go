package receivers

// TODO(gri) add more tests checking the various restrictions on receivers

type G[P any] struct{}
type A[P any] = G[P]

func (a A /* ERROR "cannot define new methods on generic alias type A[P any]" */ [P]) m() {}
