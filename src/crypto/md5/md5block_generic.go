//go:build (!riscv64 && !s390x) || purego

package md5

const haveAsm = false

func block(dig *digest, p []byte) {
	blockGeneric(dig, p)
}
