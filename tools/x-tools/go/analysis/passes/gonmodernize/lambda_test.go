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

func TestLambda(t *testing.T) {
	analysistest.RunWithSuggestedFixes(t, analysistest.TestData(), gonmodernize.LambdaAnalyzer, "lambda", "lambdaimports", "lambdadot")
	// Check the complete transformed fixture together: individual edits must
	// remain valid when every suggested fix is applied in the same operation.
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filepath.Join(analysistest.TestData(), "src/lambda/lambda.go.golden"), nil, 0)
	if err != nil {
		t.Fatal(err)
	}
	config := types.Config{GoVersion: "go1.27"}
	if _, err := config.Check("lambda", fset, []*ast.File{file}, nil); err != nil {
		t.Fatal(err)
	}
}
