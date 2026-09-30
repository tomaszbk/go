package bug1

import "./bug0"

// This is expected to fail--t0 is in package bug0 and should not be
// visible here in package bug1.  The test for failure is in
// ../bug083.go.

var v1 bug0.t0;	// ERROR "bug0"

