//go:build !aix && !darwin && !freebsd && !openbsd && !solaris && !wasip1 && !windows && !(linux && amd64) && !plan9

package runtime

//go:wasmimport gojs runtime.walltime
func walltime() (sec int64, nsec int32)
