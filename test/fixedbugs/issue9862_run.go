// run

//go:build !nacl && !js && !wasip1 && gc


// Check for compile or link error.

package main

import (
	"os/exec"
	"strings"
)

func main() {
	out, err := exec.Command("go", "run", "fixedbugs/issue9862.go").CombinedOutput()
	outstr := string(out)
	if err == nil {
		println("go run issue9862.go succeeded, should have failed\n", outstr)
		return
	}
	if !strings.Contains(outstr, "symbol too large") {
		println("go run issue9862.go gave unexpected error; want symbol too large:\n", outstr)
	}
}
