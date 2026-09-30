//go:build !aix && !darwin && !freebsd && !openbsd && !plan9 && !solaris && !wasip1

package runtime

//go:wasmimport gojs runtime.nanotime1
func nanotime1() int64
