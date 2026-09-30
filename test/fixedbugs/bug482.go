// compile


// Using the same name for a field in a composite literal and for a
// global variable that depends on the variable being initialized
// caused gccgo to erroneously report "variable initializer refers to
// itself".

package p

type S struct {
	F int
}

var V = S{F: 1}

var F = V.F
