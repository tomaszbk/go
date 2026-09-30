package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) > 1 {
		// Generate a SIGILL.
		sigill()
		return
	}
	// Run ourselves with an extra argument. That process should SIGILL.
	out, _ := exec.Command(os.Args[0], "foo").CombinedOutput()
	want := "instruction bytes: 0xf 0xb 0xc3"
	if !bytes.Contains(out, []byte(want)) {
		fmt.Printf("got:\n%s\nwant:\n%s\n", string(out), want)
	}
}
func sigill()
