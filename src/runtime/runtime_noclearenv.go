//go:build !linux

package runtime

import _ "unsafe" // for go:linkname

//go:linkname syscall_runtimeClearenv syscall.runtimeClearenv
func syscall_runtimeClearenv(env map[string]int) {
	// The system doesn't have clearenv(3) so emulate it by unsetting all of
	// the variables manually.
	for k := range env {
		syscall_runtimeUnsetenv(k)
	}
}
