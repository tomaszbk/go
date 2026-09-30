//go:build plan9 || windows

package doc

import (
	"os"
)

var signalsToIgnore = []os.Signal{os.Interrupt}
