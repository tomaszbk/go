//go:build test_run

package main

import "cgostdio/stdio"

func main() {
	stdio.Stdout.WriteString(stdio.Greeting + "\n")
}
