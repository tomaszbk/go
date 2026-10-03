// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gonmodernize_test

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
	"golang.org/x/tools/go/analysis/passes/gonmodernize"
)

func TestErrors(t *testing.T) {
	results := analysistest.RunWithSuggestedFixes(t, analysistest.TestData(), gonmodernize.ErrorAnalyzer, "gonerrors")
	for _, result := range results {
		fset := token.NewFileSet()
		var files []*ast.File
		for _, original := range result.Pass.Files {
			filename := result.Pass.Fset.File(original.Pos()).Name()
			content, err := os.ReadFile(filename + ".golden")
			if os.IsNotExist(err) {
				content, err = os.ReadFile(filename)
			}
			if err != nil {
				t.Fatal(err)
			}
			file, err := parser.ParseFile(fset, filename, content, parser.AllErrors)
			if err != nil {
				t.Fatal(err)
			}
			files = append(files, file)
		}
		config := types.Config{Importer: errorFixtureImports(result.Pass.Pkg.Imports())}
		if _, err := config.Check(result.Pass.Pkg.Path(), fset, files, nil); err != nil {
			t.Fatalf("fixed error checks do not type-check: %v", err)
		}
	}
}

type errorFixtureImports []*types.Package

func (imports errorFixtureImports) Import(path string) (*types.Package, error) {
	for _, pkg := range imports {
		if pkg.Path() == path {
			return pkg, nil
		}
	}
	return nil, fmt.Errorf("unexpected fixture import %s", path)
}
