// errorcheck


// Issue 4283: nil == nil can't be done as the type is unknown.

package p

func F1() bool {
	return nil == nil	// ERROR "invalid"
}

func F2() bool {
	return nil != nil	// ERROR "invalid"
}
