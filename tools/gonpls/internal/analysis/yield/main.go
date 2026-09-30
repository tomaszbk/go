//go:build ignore

// The yield command applies the yield analyzer to the specified
// packages of Go source code.
package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"
	"golang.org/x/tools/gopls/internal/analysis/yield"
)

func main() { singlechecker.Main(yield.Analyzer) }
