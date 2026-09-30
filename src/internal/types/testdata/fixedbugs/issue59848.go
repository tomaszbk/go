package p

type T struct{}
type I interface{ M() }
var _ I = T /* ERROR "missing method M" */ {} // must not crash
func (T) m() {}
