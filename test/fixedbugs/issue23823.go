// errorcheck


package p

type I1 = interface {
	I2
}

// BAD: type loop should mention I1; see also #41669
type I2 interface { // GC_ERROR "invalid recursive type: I2 refers to itself"
	I1 // GCCGO_ERROR "invalid recursive interface"
}
