//go:build !arm && !arm64 && !mips64 && !mips64le && !mips && !mipsle && !wasm

package runtime

// careful: cputicks is not guaranteed to be monotonic! In particular, we have
// noticed drift between cpus on certain os/arch combinations. See issue 8976.
func cputicks() int64
