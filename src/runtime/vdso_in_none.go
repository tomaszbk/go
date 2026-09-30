//go:build (linux && !386 && !amd64 && !arm && !arm64 && !loong64 && !mips64 && !mips64le && !ppc64 && !ppc64le && !riscv64 && !s390x) || !linux

package runtime

// A dummy version of inVDSOPage for targets that don't use a VDSO.

func inVDSOPage(pc uintptr) bool {
	return false
}
