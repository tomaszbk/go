//go:build !(linux && (amd64 || arm64 || arm64be || ppc64 || ppc64le || loong64 || s390x || riscv64))

package runtime

import _ "unsafe"

//go:linkname vgetrandom
func vgetrandom(p []byte, flags uint32) (ret int, supported bool) {
	return -1, false
}

func vgetrandomDestroy(mp *m) {}

func vgetrandomInit() {}
