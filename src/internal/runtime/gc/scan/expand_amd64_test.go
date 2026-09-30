//go:build amd64

package scan_test

import (
	"internal/runtime/gc/scan"
	"testing"
)

func TestExpandAVX512(t *testing.T) {
	if !scan.CanAVX512() {
		t.Skip("no AVX512")
	}
	testExpand(t, scan.ExpandAVX512)
}
