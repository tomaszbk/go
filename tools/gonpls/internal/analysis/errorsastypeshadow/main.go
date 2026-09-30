//go:build ignore

// The errorsastypeshadow command runs the errorsastypeshadow analyzer.
package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"
	"golang.org/x/tools/gopls/internal/analysis/errorsastypeshadow"
)

func main() { singlechecker.Main(errorsastypeshadow.Analyzer) }
