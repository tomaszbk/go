package maprange_test

import (
	"golang.org/x/tools/go/analysis/analysistest"
	"golang.org/x/tools/gopls/internal/analysis/maprange"
	"golang.org/x/tools/internal/testfiles"
	"path/filepath"
	"testing"
)

func TestBasic(t *testing.T) {
	dir := testfiles.ExtractTxtarFileToTmp(t, filepath.Join(analysistest.TestData(), "basic.txtar"))
	analysistest.RunWithSuggestedFixes(t, dir, maprange.Analyzer, "maprange")
}

func TestOld(t *testing.T) {
	dir := testfiles.ExtractTxtarFileToTmp(t, filepath.Join(analysistest.TestData(), "old.txtar"))
	analysistest.RunWithSuggestedFixes(t, dir, maprange.Analyzer, "maprange")
}
