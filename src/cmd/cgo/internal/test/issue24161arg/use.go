//go:build darwin

package issue24161arg

/*
#cgo LDFLAGS: -framework CoreFoundation
#include <CoreFoundation/CoreFoundation.h>
*/
import "C"
import "testing"

func Test(t *testing.T) {
	a := test24161array()
	C.CFArrayCreateCopy(0, a)
}
