//go:build aix || (darwin && !race) || (linux && !race) || (freebsd && !race) || netbsd || openbsd || solaris || dragonfly || zos

package unix

import (
	"unsafe"
)

const raceenabled = false

func raceAcquire(addr unsafe.Pointer) {
}

func raceReleaseMerge(addr unsafe.Pointer) {
}

func raceReadRange(addr unsafe.Pointer, len int) {
}

func raceWriteRange(addr unsafe.Pointer, len int) {
}
