//go:build (386 || arm || loong64 || riscv64) && !purego

package sha1

//go:noescape
func block(dig *digest, p []byte)
