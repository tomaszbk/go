//go:build gccgo

package build

import "runtime"

// getToolDir returns the default value of ToolDir.
func getToolDir() string {
	return envOr("GCCGOTOOLDIR", runtime.GCCGOTOOLDIR)
}
