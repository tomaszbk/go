package syntax

import (
	"testing"
)

func TestDump(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	ast, _ := ParseFile(*src_, func(err error) { t.Error(err) }, nil, CheckBranches)

	if ast != nil {
		Fdump(testOut(), ast)
	}
}
