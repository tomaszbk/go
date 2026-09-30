//go:build linux && !386 && !amd64 && !arm && !arm64 && !loong64 && !mips64 && !mips64le && !ppc64 && !ppc64le && !riscv64 && !s390x

package runtime

func vdsoauxv(tag, val uintptr) {
}
