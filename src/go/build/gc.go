//go:build gc

package build

import (
	"path/filepath"
	"runtime"
)

// getToolDir returns the default value of ToolDir.
func getToolDir() string {
	return filepath.Join(runtime.GOROOT(), "pkg/tool/"+runtime.GOOS+"_"+runtime.GOARCH)
}
