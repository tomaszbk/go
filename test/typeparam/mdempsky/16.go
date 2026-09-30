// run

// Test that type assertion panics mention the real interface type,
// not their shape type.

package main

import (
	"fmt"
	"runtime"
	"strings"
)

func main() {
	// The exact error message isn't important, but it should mention
	// `main.T`, not `go.shape.int_0`.
	if have := F[T](); !strings.Contains(have, "interface { T() main.T }") {
		fmt.Printf("FAIL: unexpected panic message: %q\n", have)
	}
}

type T int

func F[T any]() (res string) {
	defer func() {
		res = recover().(runtime.Error).Error()
	}()
	_ = interface{ T() T }(nil).(T)
	return
}
