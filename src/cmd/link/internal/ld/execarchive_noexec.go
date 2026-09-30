//go:build wasm || windows

package ld

const syscallExecSupported = false

func (ctxt *Link) execArchive(argv []string) {
	panic("should never arrive here")
}
