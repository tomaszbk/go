//go:build !(amd64 && goexperiment.simd && !boringcrypto)

package amd64_test

func testLoop2(n int, x float64) float64 {
	return 0
}
func testLoop3(n int, x float64) float64 {
	return 0
}
