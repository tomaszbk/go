// This file contains tests for the atomic checker.

package atomic

import "sync/atomic"

func AtomicTests() {
	x := uint64(1)
	x = atomic.AddUint64(&x, 1) // ERROR "direct assignment to atomic value"
}
