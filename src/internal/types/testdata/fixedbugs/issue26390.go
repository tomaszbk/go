// stand-alone test to ensure case is triggered

package issue26390

type A = T

func (t *T) m() *A { return t }

type T struct{}
