//go:build darwin || freebsd || netbsd || openbsd

package sysinfo

import "syscall"

func osCPUInfoName() string {
	cpu, _ := syscall.Sysctl("machdep.cpu.brand_string")
	return cpu
}
