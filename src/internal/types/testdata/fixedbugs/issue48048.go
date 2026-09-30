package p

type T[P any] struct{}

func (T[_]) A() {}

var _ = (T[int]).A
var _ = (*T[int]).A

var _ = (T /* ERROR "cannot use generic type" */).A
var _ = (*T /* ERROR "cannot use generic type" */).A
