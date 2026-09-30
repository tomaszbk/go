package b

import "testshared/issue44031/a"

type T int

func (T) M() {}

var i = a.ATypeWithALoooooongName(T(0))

func F() {
	i.M()
}
