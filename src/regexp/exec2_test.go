//go:build !race

package regexp

import (
	"testing"
)

// This test is excluded when running under the race detector because
// it is a very expensive test and takes too long.
func TestRE2Exhaustive(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping TestRE2Exhaustive during short test")
	}
	testRE2(t, "testdata/re2-exhaustive.txt.bz2")
}
