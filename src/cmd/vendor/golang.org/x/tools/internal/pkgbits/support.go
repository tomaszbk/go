package pkgbits

// This package is dependency-restricted; see x/tools/go/gcexportdata.TestDeps.
import "fmt"

func assert(b bool) {
	if !b {
		panic("assertion failed")
	}
}

func panicf(format string, args ...any) {
	panic(fmt.Errorf(format, args...))
}
