//go:build !faketime

package runtime

import _ "unsafe" // for go:linkname

// Do not remove or change the type signature.
// See comment in timestub.go.
//
//go:linkname time_now time.now
func time_now() (sec int64, nsec int32, mono int64) {
	var t [4]uint64
	if readtime(&t[0], 1, 4) == 4 {
		mono = int64(frombe(t[3])) // new kernel, use monotonic time
	} else {
		mono = int64(frombe(t[0])) // old kernel, fall back to unix time
	}
	sec, nsec = timesplit(frombe(t[0]))
	return
}
