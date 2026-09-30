//go:build !amd64

package flate

// shiftMask is a no-op shift mask for non-x86-64.
// The compiler will optimize it away.
const reg8SizeMask64 = 0xff
