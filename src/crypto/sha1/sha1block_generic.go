//go:build (!386 && !amd64 && !arm && !arm64 && !loong64 && !riscv64 && !s390x) || purego

package sha1

func block(dig *digest, p []byte) {
	blockGeneric(dig, p)
}
