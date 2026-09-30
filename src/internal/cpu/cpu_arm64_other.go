//go:build arm64 && !linux && !freebsd && !android && !darwin && !openbsd && !windows

package cpu

func osInit() {
	// Other operating systems do not support reading HWCap from auxiliary vector,
	// reading privileged aarch64 system registers or sysctl in user space to detect
	// CPU features at runtime.
}
