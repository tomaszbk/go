// errorcheck

package p

func f(e interface{}) {
	switch e.(type) {
	case nil, nil: // ERROR "multiple nil cases in type switch|duplicate type in switch|duplicate case nil in type switch"
	}

	switch e.(type) {
	case nil:
	case nil: // ERROR "multiple nil cases in type switch|duplicate type in switch|duplicate case nil in type switch"
	}
}
