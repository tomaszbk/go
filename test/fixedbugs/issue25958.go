// errorcheck


package p

// Verify that the "must be receive" error for "case done:" appears
// on the line of the case clause, not the line of the done declaration.

func f(done chan struct{}) {
	select {
	case done: // ERROR "must be receive|expected .*<-.* or .*=|must be send or receive|not used"
	case (chan struct{})(done): // ERROR "must be receive|expected .*<-.* or .*=|must be send or receive"
	}
}
