// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gonmodernize

import (
	"go/ast"
	"go/token"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// LambdaAnalyzer replaces function literal signatures with Gon lambda syntax
// when an identical signature is available from the surrounding context.
var LambdaAnalyzer = &analysis.Analyzer{
	Name: "gonlambda",
	Doc: `replace function literals with contextually typed Gon lambdas

The gonlambda analyzer finds function literals whose signature is supplied by a
typed variable declaration, an assignment, or a non-generic callback argument.
It replaces the signature with (parameters) => and preserves the block body,
including closure captures, comments, return statements, and defers.

Rewrites require identical function signatures. Inferred declarations, interface
targets, parenthesized literals, generic calls, named results, unnamed parameters,
signature comments, and signatures using imported types are left unchanged.
Generated files and files containing dot imports are also left unchanged.`,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      runLambda,
}

func runLambda(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	for fileCur := range inspect.Root().Children() {
		file := fileCur.Node().(*ast.File)
		if ast.IsGenerated(file) {
			continue
		}
		// Dot-imported type names need additional binding analysis when their
		// uses disappear. Decline these files rather than removing imports.
		dotImport := false
		for _, spec := range file.Imports {
			if spec.Name != nil && spec.Name.Name == "." {
				dotImport = true
			}
		}
		if dotImport {
			continue
		}
		for cur := range fileCur.Preorder((*ast.FuncLit)(nil)) {
			literal := cur.Node().(*ast.FuncLit)
			actual, ok := pass.TypesInfo.TypeOf(literal).(*types.Signature)
			if !ok || lambdaWithinGenericCall(pass.TypesInfo, cur) {
				continue
			}
			target := lambdaTarget(pass.TypesInfo, cur, literal)
			if target == nil {
				continue
			}
			if _, ok := target.Underlying().(*types.Signature); !ok || !types.Identical(target.Underlying(), actual) {
				continue
			}
			if replacement, ok := lambdaHeader(pass.TypesInfo, file, literal); ok {
				pass.Report(analysis.Diagnostic{
					Pos:     literal.Pos(),
					End:     literal.Type.End(),
					Message: "function literal can use Gon lambda syntax",
					SuggestedFixes: []analysis.SuggestedFix{{
						Message: "Convert function literal to lambda",
						// Keep the block, preserving comments, return/defer
						// behavior, and allowing independent fixes in its body.
						TextEdits: []analysis.TextEdit{{
							Pos:     literal.Pos(),
							End:     literal.Body.Pos(),
							NewText: []byte(replacement),
						}},
					}},
				})
			}
		}
	}
	return nil, nil
}

// lambdaTarget admits only contexts that pass a known function type directly
// to the expression. Parentheses, inferred declarations and interface targets
// do not supply that context.
func lambdaTarget(info *types.Info, cur inspector.Cursor, literal *ast.FuncLit) types.Type {
	switch parent := cur.Parent().Node().(type) {
	case *ast.ValueSpec:
		if parent.Type != nil {
			return info.TypeOf(parent.Type)
		}
	case *ast.AssignStmt:
		if parent.Tok == token.ASSIGN && len(parent.Lhs) == len(parent.Rhs) {
			for i, rhs := range parent.Rhs {
				if rhs == literal {
					return info.TypeOf(parent.Lhs[i])
				}
			}
		}
	case *ast.CallExpr:
		if info.Types[parent.Fun].IsType() || parent.Ellipsis.IsValid() {
			return nil
		}
		callee := info.TypeOf(parent.Fun)
		if callee == nil {
			return nil
		}
		sig, ok := callee.Underlying().(*types.Signature)
		if !ok || sig.TypeParams().Len() != 0 {
			return nil
		}
		for i, arg := range parent.Args {
			if arg != literal {
				continue
			}
			if sig.Variadic() && i >= sig.Params().Len()-1 {
				return sig.Params().At(sig.Params().Len() - 1).Type().(*types.Slice).Elem()
			}
			if i < sig.Params().Len() {
				return sig.Params().At(i).Type()
			}
		}
	}
	return nil
}

func lambdaWithinGenericCall(info *types.Info, cur inspector.Cursor) bool {
	for enclosing := range cur.Enclosing() {
		call, ok := enclosing.Node().(*ast.CallExpr)
		if !ok {
			continue
		}
		fun := ast.Unparen(call.Fun)
		for {
			switch index := fun.(type) {
			case *ast.IndexExpr:
				fun = ast.Unparen(index.X)
			case *ast.IndexListExpr:
				fun = ast.Unparen(index.X)
			default:
				goto unwrapped
			}
		}
	unwrapped:
		var id *ast.Ident
		switch f := fun.(type) {
		case *ast.Ident:
			id = f
		case *ast.SelectorExpr:
			id = f.Sel
		}
		if id != nil && info.Instances[id].TypeArgs != nil {
			return true
		}
	}
	return false
}

func lambdaHeader(info *types.Info, file *ast.File, literal *ast.FuncLit) (string, bool) {
	if literal.Type.Results != nil {
		for _, field := range literal.Type.Results.List {
			if len(field.Names) != 0 {
				return "", false
			}
		}
	}
	// Do not remove any package references: several individually safe fixes
	// could otherwise collectively remove an import's final uses.
	for node := range ast.Preorder(literal.Type) {
		if id, ok := node.(*ast.Ident); ok {
			if _, ok := info.Uses[id].(*types.PkgName); ok {
				return "", false
			}
		}
	}
	for _, group := range file.Comments {
		if literal.Pos() <= group.Pos() && group.Pos() < literal.Body.Pos() {
			return "", false
		}
	}
	var names []string
	for _, field := range literal.Type.Params.List {
		if len(field.Names) == 0 {
			return "", false
		}
		for _, name := range field.Names {
			names = append(names, name.Name)
		}
	}
	return "(" + strings.Join(names, ", ") + ") => ", true
}
