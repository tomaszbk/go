//go:build wasm

package drbg

func getEntropy() *[SeedSize]byte {
	panic("FIPS 140-3 entropy generation is not supported on Wasm")
}
