package unusedfunc_test

import (
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
	"golang.org/x/tools/gopls/internal/analysis/unusedfunc"
	"golang.org/x/tools/internal/testfiles"
)

func Test(t *testing.T) {
	dir := testfiles.ExtractTxtarFileToTmp(t, filepath.Join(analysistest.TestData(), "basic.txtar"))
	analysistest.RunWithSuggestedFixes(t, dir, unusedfunc.Analyzer, "example.com/a")
}
