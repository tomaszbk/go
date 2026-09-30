package fs_test

import (
	"fmt"
	. "io/fs"
	"testing"
)

type statOnly struct{ StatFS }

func (statOnly) Open(name string) (File, error) { return nil, ErrNotExist }

func TestStat(t *testing.T) {
	check := func(desc string, info FileInfo, err error) {
		t.Helper()
		if err != nil || info == nil || info.Mode() != 0456 {
			infoStr := "<nil>"
			if info != nil {
				infoStr = fmt.Sprintf("FileInfo(Mode: %#o)", info.Mode())
			}
			t.Fatalf("Stat(%s) = %v, %v, want Mode:0456, nil", desc, infoStr, err)
		}
	}

	// Test that Stat uses the method when present.
	info, err := Stat(statOnly{testFsys}, "hello.txt")
	check("statOnly", info, err)

	// Test that Stat uses Open when the method is not present.
	info, err = Stat(openOnly{testFsys}, "hello.txt")
	check("openOnly", info, err)
}
