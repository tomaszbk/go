package pkgbits

import "fmt"

func assert(b bool) {
	if !b {
		panic("assertion failed")
	}
}

func panicf(format string, args ...any) {
	panic(fmt.Errorf(format, args...))
}
