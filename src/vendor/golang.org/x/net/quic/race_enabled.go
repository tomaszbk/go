//go:build race

package quic

import (
	"runtime"
	"unsafe"
)

// Our synchronization here is rather coarse-grained since we always use the
// address of quicSync. As such, it will not differentiate between, for
// example, reads and writes to different QUIC streams.
// However, this is consistent with our synchronization for net/http, and
// mostly only matters for testing.
var quicSync uint64

func raceAcquire() {
	runtime.RaceAcquire(unsafe.Pointer(&quicSync))
}

func raceReleaseMerge() {
	runtime.RaceReleaseMerge(unsafe.Pointer(&quicSync))
}
