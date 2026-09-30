//go:build windows && !race

package windows

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
