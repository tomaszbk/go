// The unusedparams command runs the unusedparams analyzer.
package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"
	"golang.org/x/tools/gopls/internal/analysis/unusedparams"
)

func main() { singlechecker.Main(unusedparams.Analyzer) }
