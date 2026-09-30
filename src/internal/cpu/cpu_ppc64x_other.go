//go:build (ppc64 || ppc64le) && !aix && !linux

package cpu

func osinit() {
	// Other operating systems do not support reading HWCap from auxiliary vector,
	// reading privileged system registers or sysctl in user space to detect CPU
	// features at runtime.
}
