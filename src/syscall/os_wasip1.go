package syscall

//go:wasmimport wasi_snapshot_preview1 proc_exit
func ProcExit(code int32)
