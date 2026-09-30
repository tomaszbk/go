//go:build (386 || loong64 || riscv64) && !purego

package sha256

//go:noescape
func block(dig *Digest, p []byte)
