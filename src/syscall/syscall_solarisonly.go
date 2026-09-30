//go:build solaris && !illumos

package syscall

// F_DUP2FD_CLOEXEC has different values on Solaris and Illumos.
const F_DUP2FD_CLOEXEC = 0x30
