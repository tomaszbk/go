package sharedcheck_test

import (
	"golang.org/x/tools/go/analysis/analysistest"
	"honnef.co/go/tools/internal/sharedcheck"
	"testing"
)

func TestGonTargetTypes(t *testing.T) {
	a := sharedcheck.RedundantTypeInDeclarationChecker("should", true)
	a.Name = "gontargettypes"
	a.Doc = "preserve contextual target types"
	analysistest.Run(t, analysistest.TestData(), a, "gontargettypes")
}
