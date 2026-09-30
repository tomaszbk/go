//go:build wasm

package cpu

// We're compiling the cpu package for an unknown (software-abstracted) CPU.
// Make CacheLinePad an empty struct and hope that the usual struct alignment
// rules are good enough.

const cacheLineSize = 0

func initOptions() {}

func archInit() {}
