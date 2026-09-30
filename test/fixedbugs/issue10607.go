// run

//go:build linux && !ppc64 && gc && cgo

// Test that a -B option is passed through when using both internal
// and external linking mode.

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	test("internal")
	test("external") // The 'cgo' build constraint should imply that a linker is available.
}

func test(linkmode string) {
	out, err := exec.Command("go", "run", "-ldflags", "-B=0x12345678 -linkmode="+linkmode, filepath.Join("fixedbugs", "issue10607a.go")).CombinedOutput()
	if err != nil {
		fmt.Printf("BUG: linkmode=%s %v\n%s\n", linkmode, err, out)
		os.Exit(1)
	}
}
