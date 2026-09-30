//go:build (loong64 || riscv64) && !purego

package sha512

//go:noescape
func block(dig *Digest, p []byte)
