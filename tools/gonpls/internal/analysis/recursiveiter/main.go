//go:build ignore

// The recursiveiter command applies the yield analyzer to the
// specified packages of Go source code.
package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"
	"golang.org/x/tools/gopls/internal/analysis/recursiveiter"
)

func main() { singlechecker.Main(recursiveiter.Analyzer) }
