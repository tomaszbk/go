package runtime

import (
	"internal/runtime/sys"
	_ "unsafe"
)

//go:linkname dit_setEnabled crypto/subtle.setDITEnabled
func dit_setEnabled() bool {
	g := getg()
	g.ditWanted = true
	g.m.ditEnabled = true
	return sys.EnableDIT()
}

//go:linkname dit_setDisabled crypto/subtle.setDITDisabled
func dit_setDisabled() {
	g := getg()
	g.ditWanted = false
	g.m.ditEnabled = false
	sys.DisableDIT()
}
