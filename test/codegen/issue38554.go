// asmcheck

// Test that we are zeroing directly instead of
// copying a large zero value. Issue 38554.

package codegen

func retlarge() [256]byte {
	// amd64:-"DUFFCOPY"
	return [256]byte{}
}
