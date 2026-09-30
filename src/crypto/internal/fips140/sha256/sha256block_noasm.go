//go:build (!386 && !amd64 && !arm64 && !loong64 && !ppc64 && !ppc64le && !riscv64 && !s390x) || purego

package sha256

func block(dig *Digest, p []byte) {
	blockGeneric(dig, p)
}
