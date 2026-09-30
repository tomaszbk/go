//go:build (linux && 386) || (linux && arm) || (linux && mips) || (linux && mipsle)

package syscall

func init() {
	// On 32-bit Linux systems, the fcntl syscall that matches Go's
	// Flock_t type is SYS_FCNTL64, not SYS_FCNTL.
	fcntl64Syscall = SYS_FCNTL64
}
