// errorcheck

// Issue 6405: spurious 'not enough arguments to return' error

package p

func Open() (int, error) {
	return OpenFile() // ERROR "undefined: OpenFile"
}
