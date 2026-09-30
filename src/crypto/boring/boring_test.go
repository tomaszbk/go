//go:build boringcrypto

package boring_test

import (
	"crypto/boring"
	"runtime"
	"testing"
)

func TestEnabled(t *testing.T) {
	supportedPlatform := runtime.GOOS == "linux" && (runtime.GOARCH == "amd64" || runtime.GOARCH == "arm64")
	if supportedPlatform && !boring.Enabled() {
		t.Error("Enabled returned false on a supported platform")
	} else if !supportedPlatform && boring.Enabled() {
		t.Error("Enabled returned true on an unsupported platform")
	}
}
