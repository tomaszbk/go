// errorcheck

package p

type T interface{ M() }

func F() T

var _ = F().(*X) // ERROR "undefined: X"
