//go:build amd64 || arm64 || loong64 || s390x || ppc64le || ppc64

package bytealg

// Index returns the index of the first instance of b in a, or -1 if b is not present in a.
// Requires 2 <= len(b) <= MaxLen.
//
//go:noescape
func Index(a, b []byte) int

// IndexString returns the index of the first instance of b in a, or -1 if b is not present in a.
// Requires 2 <= len(b) <= MaxLen.
//
//go:noescape
func IndexString(a, b string) int
