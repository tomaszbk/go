// compile

package p

type T[U interface{ M() T[U] }] int

type X int

func (X) M() T[X] { return 0 }
