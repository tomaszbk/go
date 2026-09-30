package scan_test

import (
	"internal/runtime/gc/scan"
	"testing"
)

func TestScanSpanPackedGo(t *testing.T) {
	testScanSpanPacked(t, scan.ScanSpanPackedGo)
}
