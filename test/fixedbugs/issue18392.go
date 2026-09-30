// errorcheck

package p

type A interface { // ERROR "invalid recursive type"
	Fn(A.Fn)
}
