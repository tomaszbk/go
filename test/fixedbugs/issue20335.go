// compile


// Issue 20335: don't reorder loads with stores.
// This test should fail on the ssacheck builder
// without the fix in the CL that added this file.
// TODO: check the generated assembly?

package a

import "sync/atomic"

func f(p, q *int32) bool {
	x := *q
	return atomic.AddInt32(p, 1) == x
}
