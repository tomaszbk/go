package copylock_test

import (
	"golang.org/x/tools/go/analysis/analysistest"
	"golang.org/x/tools/go/analysis/passes/copylock"
	"testing"
)

func TestGonFeatures(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), copylock.Analyzer, "gonfeatures")
}
