//go:build arm

// Make sure we can compile assembly with DIV and MOD in it.
// They get rewritten to runtime calls on GOARM=5.

package main

func f(x, y uint32)

func main() {
	f(5, 8)
}
