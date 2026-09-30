package modfetch

import (
	"context"
	"path/filepath"
	"testing"
)

func TestWriteDiskCache(t *testing.T) {
	ctx := context.Background()

	tmpdir := t.TempDir()
	err := writeDiskCache(ctx, filepath.Join(tmpdir, "file"), []byte("data"))
	if err != nil {
		t.Fatal(err)
	}
}
