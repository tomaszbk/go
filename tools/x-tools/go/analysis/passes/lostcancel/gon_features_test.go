package lostcancel_test

import (
	"golang.org/x/tools/go/analysis/analysistest"
	"golang.org/x/tools/go/analysis/passes/lostcancel"
	"testing"
)

func TestGonLambdaBoundary(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), lostcancel.Analyzer, "gonlambda")
}
