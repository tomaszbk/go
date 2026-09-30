//go:build !unix || wasm

package tls

func pauseProcess() {
	panic("-wait-for-debugger not supported on this OS")
}
