//go:build unix && !darwin

package syscall

// adjustFileLimit adds per-OS limitations on the Rlimit used for RLIMIT_NOFILE. See rlimit.go.
func adjustFileLimit(lim *Rlimit) {}
