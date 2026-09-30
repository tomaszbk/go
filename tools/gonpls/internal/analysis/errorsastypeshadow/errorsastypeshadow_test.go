package errorsastypeshadow_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
	"golang.org/x/tools/gopls/internal/analysis/errorsastypeshadow"
	"golang.org/x/tools/internal/testenv"
)

func Test(t *testing.T) {
	testenv.NeedsGo1Point(t, 26) // AsType introduced in 1.26
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, errorsastypeshadow.Analyzer, "errorsastypeshadow")
}
