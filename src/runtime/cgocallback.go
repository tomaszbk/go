package runtime

// These functions are called from C code via cgo/callbacks.go.

// Panic.

func _cgo_panic_internal(p *byte) {
	panic(gostringnocopy(p))
}
