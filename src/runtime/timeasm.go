// Declarations for operating systems implementing time.now directly in assembly.

//go:build !faketime && (windows || (linux && amd64))

package runtime

import _ "unsafe"

//go:linkname time_now time.now
func time_now() (sec int64, nsec int32, mono int64)
