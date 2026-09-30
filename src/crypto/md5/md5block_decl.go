//go:build (riscv64 || s390x) && !purego

package md5

const haveAsm = true

//go:noescape
func block(dig *digest, p []byte)
