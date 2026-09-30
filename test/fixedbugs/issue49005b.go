// errorcheck


package p

type T interface{ M() }

func F() T

var _ = F().(*X) // ERROR "impossible type assertion:( F\(\).\(\*X\))?\n\t\*X does not implement T \(missing method M\)"

type X struct{}
