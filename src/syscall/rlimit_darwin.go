//go:build darwin

package syscall

// adjustFileLimit adds per-OS limitations on the Rlimit used for RLIMIT_NOFILE. See rlimit.go.
func adjustFileLimit(lim *Rlimit) {
	// On older macOS, setrlimit(RLIMIT_NOFILE, lim) with lim.Cur = infinity fails.
	// Set to the value of kern.maxfilesperproc instead.
	n, err := SysctlUint32("kern.maxfilesperproc")
	if err != nil {
		return
	}
	if lim.Cur > uint64(n) {
		lim.Cur = uint64(n)
	}
}
