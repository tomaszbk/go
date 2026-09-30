package unusedvariable_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
	"golang.org/x/tools/gopls/internal/analysis/unusedvariable"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()

	t.Run("decl", func(t *testing.T) {
		analysistest.RunWithSuggestedFixes(t, testdata, unusedvariable.Analyzer, "decl")
	})

	t.Run("assign", func(t *testing.T) {
		analysistest.RunWithSuggestedFixes(t, testdata, unusedvariable.Analyzer, "assign")
	})
}
