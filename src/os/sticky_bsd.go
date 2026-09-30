//go:build aix || darwin || dragonfly || freebsd || (js && wasm) || netbsd || openbsd || solaris || wasip1

package os

// According to sticky(8), neither open(2) nor mkdir(2) will create
// a file with the sticky bit set.
const supportsCreateWithStickyBit = false
