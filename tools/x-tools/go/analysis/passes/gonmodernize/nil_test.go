// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gonmodernize_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
	"golang.org/x/tools/go/analysis/passes/gonmodernize"
)

func TestNil(t *testing.T) {
	analysistest.RunWithSuggestedFixes(t, analysistest.TestData(), gonmodernize.NilAnalyzer, "nil")
	checkConditionalGolden(t, "nil", "nil.go")
}
