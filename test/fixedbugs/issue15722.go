// errorcheck


// Checks to make sure that the compiler can catch a specific invalid
// method type expression. NB: gccgo and gc have slightly different
// error messages, hence the generic test for 'method' and not something
// more specific.

package issue15722

type T int
type P *T

func (T) t() {}

func _(p P) {
	P.t(p) // ERROR "method"
}
