package os

// Hostname returns the host name reported by the kernel.
func Hostname() (name string, err error) {
	return hostname()
}
