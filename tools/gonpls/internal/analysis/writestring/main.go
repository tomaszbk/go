//go:build ignore

// The writestring command applies the golang.org/x/tools/go/analysis/passes/writestring
// analysis to the specified packages of Go source code.
package main

import (
	"golang.org/x/tools/go/analysis/passes/writestring"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(writestring.Analyzer) }
