//go:build aix || linux || solaris

package osinfo

import (
	"bytes"
	"strings"
	"unsafe"
)

// Version returns the OS version name/number.
func Version() (string, error) {
	var uts utsname
	if err := uname(&uts); err != nil {
		return "", err
	}

	var sb strings.Builder

	writeCStr := func(b []byte) {
		if i := bytes.IndexByte(b, '\000'); i >= 0 {
			b = b[:i]
		}
		sb.Write(b)
	}

	// We need some absurd conversions because syscall.Utsname
	// sometimes uses []uint8 and sometimes []int8.

	s := uts.Sysname[:]
	writeCStr(*(*[]byte)(unsafe.Pointer(&s)))
	sb.WriteByte(' ')
	s = uts.Release[:]
	writeCStr(*(*[]byte)(unsafe.Pointer(&s)))
	sb.WriteByte(' ')
	s = uts.Version[:]
	writeCStr(*(*[]byte)(unsafe.Pointer(&s)))
	sb.WriteByte(' ')
	s = uts.Machine[:]
	writeCStr(*(*[]byte)(unsafe.Pointer(&s)))

	return sb.String(), nil
}
