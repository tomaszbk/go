//go:build ignore

// The unusedfunc command runs the maprange analyzer.
package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"
	"golang.org/x/tools/gopls/internal/analysis/maprange"
)

func main() { singlechecker.Main(maprange.Analyzer) }
