//go:build unix

package osinfo

import (
	"strings"
	"testing"
)

func TestVersion(t *testing.T) {
	v, err := Version()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%q", v)

	fields := strings.Fields(v)
	if len(fields) < 4 {
		t.Errorf("wanted at least 4 fields in %q, got %d", v, len(fields))
	}
}
