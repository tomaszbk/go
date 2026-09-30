package writestring_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
	"golang.org/x/tools/gopls/internal/analysis/writestring"
)

func Test(t *testing.T) {
	analysistest.RunWithSuggestedFixes(t, analysistest.TestData(), writestring.Analyzer, "a")
}
