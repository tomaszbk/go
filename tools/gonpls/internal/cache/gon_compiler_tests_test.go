package cache

import (
	"path/filepath"
	"testing"
)

func TestGonCompilerTest(t *testing.T) {
	root := t.TempDir()
	for _, tt := range []struct {
		name, path, source string
		want               bool
	}{
		{"run", "test/235.go", "// run\n\npackage main", true},
		{"flags", "test/fixedbugs/issue.go", "// errorcheck -0 -m\npackage p", true},
		{"constraints", "test/noinit.go", "//go:build linux\n\n// run\npackage main", true},
		{"directory driver", "test/issue.go", "// rundir\npackage ignored", true},
		{"skipped input", "test/skipped.go", "// skip\npackage main", true},
		{"fixture", "test/issue.dir/main.go", "// run\npackage main", false},
		{"ordinary package", "src/example/main.go", "// run\npackage main", false},
		{"outside tree", "testdata/main.go", "// run\npackage main", false},
		{"no recipe", "test/ordinary.go", "package main\n// run", false},
		{"unknown recipe", "test/ordinary.go", "// runtime example\npackage main", false},
		{"later comment", "test/ordinary.go", "// documentation\n// run\npackage main", false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCompilerTest(root, filepath.Join(root, filepath.FromSlash(tt.path)), []byte(tt.source)); got != tt.want {
				t.Errorf("isCompilerTest = %v, want %v", got, tt.want)
			}
		})
	}
	if isCompilerTest("", filepath.Join(root, "test", "235.go"), []byte("// run\npackage main")) {
		t.Fatal("empty GOROOT must not match")
	}
}
