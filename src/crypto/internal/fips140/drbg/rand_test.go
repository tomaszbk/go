package drbg

import (
	"crypto/internal/fips140"
	"testing"
)

func BenchmarkDBRG(b *testing.B) {
	old := fips140.Enabled
	defer func() {
		fips140.Enabled = old
	}()
	fips140.Enabled = true

	const N = 64
	b.SetBytes(N)
	b.RunParallel(func(pb *testing.PB) {
		buf := make([]byte, N)
		for pb.Next() {
			Read(buf)
		}
	})
}
