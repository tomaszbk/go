// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gonmodernize_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
	"golang.org/x/tools/go/analysis/passes/gonmodernize"
)

func TestConditional(t *testing.T) {
	analysistest.RunWithSuggestedFixes(t, analysistest.TestData(), gonmodernize.ConditionalAnalyzer, "conditional")
	checkConditionalGolden(t, "conditional", "conditional.go")
}

func checkConditionalGolden(t *testing.T, pkg, name string) {
	t.Helper()
	fset := token.NewFileSet()
	path := filepath.Join(analysistest.TestData(), "src", pkg, name+".golden")
	file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		t.Fatal(err)
	}
	var config types.Config
	if _, err := config.Check(pkg, fset, []*ast.File{file}, nil); err != nil {
		t.Fatalf("suggested fix must type check: %v", err)
	}
}
