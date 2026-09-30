//go:build fips140v1.0

package rand

import "io"

func fips140SetTestingReader(r io.Reader) {
	panic("cryptotest.SetGlobalRandom is not supported when building against Go Cryptographic Module v1.0.0")
}
