package fdtest

import (
	"os"
	"runtime"
	"testing"
)

func TestExists(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Exists not implemented for windows")
	}

	if !Exists(os.Stdout.Fd()) {
		t.Errorf("Exists(%d) got false want true", os.Stdout.Fd())
	}
}
