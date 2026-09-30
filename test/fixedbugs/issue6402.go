// errorcheck


// Issue 6402: spurious 'use of untyped nil' error

package p

func f() uintptr {
	return nil // ERROR "cannot use nil as type uintptr in return argument|incompatible type|cannot use nil"
}
