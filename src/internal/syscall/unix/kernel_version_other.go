//go:build !freebsd && !linux && !solaris

package unix

func KernelVersion() (major int, minor int) {
	return 0, 0
}
