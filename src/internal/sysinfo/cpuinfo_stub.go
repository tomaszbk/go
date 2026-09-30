//go:build !(darwin || freebsd || linux || netbsd || openbsd)

package sysinfo

func osCPUInfoName() string {
	return ""
}
