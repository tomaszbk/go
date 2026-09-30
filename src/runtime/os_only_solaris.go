// Solaris code that doesn't also apply to illumos.

//go:build !illumos

package runtime

func getCPUCount() int32 {
	n := int32(sysconf(__SC_NPROCESSORS_ONLN))
	if n < 1 {
		return 1
	}

	return n
}
