// Copyright 2026 The Gon Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gonmodernize_test

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
	"testing"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/checker"
	"golang.org/x/tools/go/analysis/passes/gonmodernize"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/internal/analysis/driverutil"
)

// TestSuiteCombinedFixes uses the same edit merger as gon fix. A whole-handler,
// whole-branch, or whole-call edit can overlap another analyzer's smaller edit.
// All orders must leave valid source after every round, including when the
// merger skips a conflicting fix and asks the user to rerun the command.
// Executable legacy/modern behavior pairs live in misc/gon/test_fix.py.
func TestSuiteCombinedFixes(t *testing.T) {
	var orders func([]*analysis.Analyzer, int)
	orders = func(analyzers []*analysis.Analyzer, start int) {
		if start < len(analyzers) {
			for i := start; i < len(analyzers); i++ {
				analyzers[start], analyzers[i] = analyzers[i], analyzers[start]
				orders(analyzers, start+1)
				analyzers[start], analyzers[i] = analyzers[i], analyzers[start]
			}
			return
		}
		var names []string
		for _, analyzer := range analyzers {
			names = append(names, analyzer.Name)
		}
		t.Run(strings.Join(names, "_"), func(t *testing.T) {
			filename := filepath.Join(t.TempDir(), "combined.go")
			source := []byte(suiteCombinedSource)
			for round := range 8 {
				if err := os.WriteFile(filename, source, 0600); err != nil {
					t.Fatal(err)
				}
				pkg := suitePackage(t, filename, source)
				graph, err := checker.Analyze(analyzers, []*packages.Package{pkg}, nil)
				if err != nil {
					t.Fatal(err)
				}
				var actions []driverutil.FixAction
				count := 0
				for _, action := range graph.Roots {
					if action.Err != nil {
						t.Fatal(action.Err)
					}
					if round == 0 && len(action.Diagnostics) == 0 {
						t.Fatalf("fixture does not exercise %s", action.Analyzer.Name)
					}
					count += len(action.Diagnostics)
					actions = append(actions, driverutil.FixAction{
						Name: action.String(), Pkg: pkg.Types, Files: pkg.Syntax,
						FileSet: pkg.Fset, ReadFileFunc: os.ReadFile, Diagnostics: action.Diagnostics,
					})
				}
				if count == 0 {
					for _, syntax := range []string{" or err ", "!", " = if ", "??=", "?(", "=>"} {
						if !bytes.Contains(source, []byte(syntax)) {
							t.Errorf("combined fixes did not exercise %q:\n%s", syntax, source)
						}
					}
					return
				}
				var updated []byte
				err = driverutil.ApplyFixes(actions, func(name string, content []byte) error {
					if name != filename {
						t.Fatalf("unexpected file update %q", name)
					}
					updated = content
					return nil
				}, false, false)
				if err != nil && !strings.Contains(err.Error(), "Re-run the command to apply more") {
					t.Fatal(err)
				}
				if len(updated) == 0 || bytes.Equal(updated, source) {
					t.Fatalf("round %d made no progress with %d diagnostics: %v", round, count, err)
				}
				// In particular, lambdas moved into a conditional branch or an
				// optional call must still receive the identical contextual type.
				suitePackage(t, filename, updated)
				source = updated
			}
			t.Fatal("combined fixes did not converge")
		})
	}
	orders(slices.Clone(gonmodernize.Suite), 0)
}

func suitePackage(t *testing.T, filename string, source []byte) *packages.Package {
	t.Helper()
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, source, parser.ParseComments)
	if err != nil {
		t.Fatalf("parse combined fixes: %v\n%s", err, source)
	}
	info := &types.Info{
		Types:      make(map[ast.Expr]types.TypeAndValue),
		Defs:       make(map[*ast.Ident]types.Object),
		Uses:       make(map[*ast.Ident]types.Object),
		Instances:  make(map[*ast.Ident]types.Instance),
		Implicits:  make(map[ast.Node]types.Object),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
		Scopes:     make(map[ast.Node]*types.Scope),
	}
	sizes := types.SizesFor("gc", runtime.GOARCH)
	config := types.Config{GoVersion: "go1.27", Sizes: sizes}
	pkg, err := config.Check("combined", fset, []*ast.File{file}, info)
	if err != nil {
		t.Fatalf("typecheck combined fixes: %v\n%s", err, source)
	}
	return &packages.Package{
		ID: "combined", Name: "combined", PkgPath: "combined",
		GoFiles: []string{filename}, CompiledGoFiles: []string{filename},
		Fset: fset, Syntax: []*ast.File{file}, Types: pkg, TypesInfo: info, TypesSizes: sizes,
	}
}

const suiteCombinedSource = `package combined

type failure string
func (f failure) Error() string { return string(f) }

func load(failed bool) (int, error) {
	if failed { return 0, failure("failed") }
	return 7, nil
}

func propagate(failed bool) (int, error) {
	value, err := load(failed)
	if err != nil { return 0, err }
	return value, nil
}

func nested(failed, choose bool, ptr *int) (int, error) {
	value, err := load(failed)
	if err != nil {
		var handler func(int) int
		if choose {
			handler = func(x int) int { return x + 1 }
		} else {
			handler = func(x int) int { return x + 2 }
		}
		if ptr == nil { ptr = new(int) }
		return handler(*ptr), err
	}
	var adjust func(int) int = func(x int) int {
		if choose { return x + 3 } else { return x + 4 }
	}
	var selected func(int) int
	if choose {
		selected = func(x int) int { return x + 5 }
	} else {
		selected = func(x int) int { return x + 6 }
	}
	var callback func(func(int) int) = func(f func(int) int) { value = f(value) }
	if choose { callback = nil }
	if callback != nil { callback(func(x int) int { return x + 7 }) }
	if ptr == nil { ptr = new(int) }
	return selected(adjust(value)) + *ptr, nil
}
`
