//go:build unix

package syscall

import _ "unsafe" // for linkname

// mmap should be an internal detail,
// but widely used packages access it using linkname.
// Notable members of the hall of shame include:
//   - modernc.org/memory
//   - github.com/ncruces/go-sqlite3
//
// Do not remove or change the type signature.
// See go.dev/issue/67401.
//
//go:linkname mmap
