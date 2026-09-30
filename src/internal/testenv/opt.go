//go:build !noopt

package testenv

// OptimizationOff reports whether optimization is disabled.
func OptimizationOff() bool {
	return false
}
