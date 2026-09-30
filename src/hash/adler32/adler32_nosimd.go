//go:build !((amd64 || arm64) && goexperiment.simd)

package adler32

const (
	haveSIMD = false
	minSIMD  = 0
)

func updateSIMD(d digest, p []byte) digest {
	panic("unreachable")
}
