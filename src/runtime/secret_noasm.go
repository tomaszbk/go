//go:build !arm64 && !amd64 && !loong64

package runtime

func secretEraseRegisters() {
	throw("runtime/secret.Do not supported yet")
}
