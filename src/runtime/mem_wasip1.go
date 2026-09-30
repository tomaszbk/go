//go:build wasip1

package runtime

func resetMemoryDataView() {
	// This function is a no-op on WASI, it is only used to notify the browser
	// that its view of the WASM memory needs to be updated when compiling for
	// GOOS=js.
}
