//go:build !arm64

package sys

var DITSupported = false

func EnableDIT() bool  { return false }
func DITEnabled() bool { return false }
func DisableDIT()      {}
