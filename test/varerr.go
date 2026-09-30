// errorcheck


// Verify that a couple of illegal variable declarations are caught by the compiler.
// Does not compile.

package main

func main() {
	_ = asdf	// ERROR "undefined.*asdf"

	new = 1	// ERROR "use of builtin new not in function call|invalid left hand side|must be called"
}

