//go:build ignore

// The unusedfunc command runs the unusedfunc analyzer.
package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"
	"golang.org/x/tools/gopls/internal/analysis/unusedfunc"
)

func main() { singlechecker.Main(unusedfunc.Analyzer) }
