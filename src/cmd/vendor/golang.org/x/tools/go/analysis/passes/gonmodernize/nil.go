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
)

// NilAnalyzer suggests nil coalescing, coalescing assignments and safe access.
var NilAnalyzer = &analysis.Analyzer{
	Name: "gonnil",
	Doc: `replace simple nil checks with Gon nil-safety operators

Suggests ??= for local lazy initialization, ?? for if/else defaults, and
?. or ?( for guarded fields and calls. The guard must test a local variable
against the predeclared nil. Transformations preserve lazy arguments and
defaults, typed-nil interfaces and original assignment/return target types.
Comments, complex locations, generic guards, nullable-result defaults and
potential changes to interface boxing are conservatively left unchanged.`,
	Run: runNil,
}

func runNil(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		if ast.IsGenerated(file) {
			continue
		}
		conditionalWalk(pass, file, nil, func(stmt *ast.IfStmt, sig *types.Signature) bool {
			if conditionalHasComments(file, stmt) {
				return false
			}
			replacement, ok := nilOpportunity(pass, stmt, sig)
			if ok {
				conditionalReport(pass, stmt, "nil check can use Gon nil-safety operators", "Use nil-safety operators", replacement)
			}
			return ok
		})
	}
	return nil, nil
}

func nilOpportunity(pass *analysis.Pass, stmt *ast.IfStmt, sig *types.Signature) (string, bool) {
	if stmt.Init != nil || len(stmt.Body.List) != 1 {
		return "", false
	}
	guard, equal := nilCondition(pass, stmt.Cond)
	if guard == nil {
		return "", false
	}
	if stmt.Else == nil {
		if assign, ok := stmt.Body.List[0].(*ast.AssignStmt); equal && ok &&
			assign.Tok == token.ASSIGN && len(assign.Lhs) == 1 && len(assign.Rhs) == 1 &&
			conditionalLocal(pass, assign.Lhs[0]) == pass.TypesInfo.ObjectOf(guard) {
			return guard.Name + " ??= " + conditionalText(pass, assign.Rhs[0]), true
		}
		if expr, ok := stmt.Body.List[0].(*ast.ExprStmt); !equal && ok {
			if _, call := ast.Unparen(expr.X).(*ast.CallExpr); call {
				return nilSafeAccess(pass, expr.X, guard)
			}
		}
		return "", false
	}
	choice, ok := conditionalChoice(pass, stmt, sig)
	if !ok {
		return "", false
	}
	present, absent := choice.then, choice.otherwise
	if equal {
		present, absent = absent, present
	}
	if nilSameVariable(pass, present, guard) {
		return choice.prefix + guard.Name + " ?? " + nilOperand(pass, absent), true
	}
	// A safe result that can itself be nil must not acquire a default: the
	// old condition only tested the receiver/callback. Even a nil default is
	// unsafe for an interface target (a typed nil could become a nonnil box).
	presentType := pass.TypesInfo.TypeOf(present)
	if presentType == nil || !types.Identical(presentType, choice.target) {
		return "", false
	}
	if _, parameter := types.Unalias(presentType).(*types.TypeParam); parameter {
		return "", false
	}
	if nilCanBeNil(presentType) {
		if !nilBuiltin(pass, absent) {
			return "", false
		}
		if safe, ok := nilSafeAccess(pass, present, guard); ok {
			return choice.prefix + safe, true
		}
		return "", false
	}
	if star, ok := ast.Unparen(present).(*ast.StarExpr); ok && nilSameVariable(pass, star.X, guard) {
		return choice.prefix + "*" + guard.Name + " ?? " + nilOperand(pass, absent), true
	}
	if safe, ok := nilSafeAccess(pass, present, guard); ok {
		return choice.prefix + safe + " ?? " + nilOperand(pass, absent), true
	}
	return "", false
}

// nilCondition only accepts local identifiers and the predeclared nil.
// Arbitrary selector/index expressions may change between the test and use.
func nilCondition(pass *analysis.Pass, expr ast.Expr) (*ast.Ident, bool) {
	condition, ok := ast.Unparen(expr).(*ast.BinaryExpr)
	if !ok || condition.Op != token.EQL && condition.Op != token.NEQ {
		return nil, false
	}
	x, y := condition.X, condition.Y
	if nilBuiltin(pass, x) {
		x, y = y, x
	}
	if !nilBuiltin(pass, y) || conditionalLocal(pass, x) == nil {
		return nil, false
	}
	id, ok := ast.Unparen(x).(*ast.Ident)
	if !ok {
		return nil, false
	}
	if _, parameter := types.Unalias(pass.TypesInfo.TypeOf(id)).(*types.TypeParam); parameter {
		return nil, false
	}
	return id, condition.Op == token.EQL
}

func nilBuiltin(pass *analysis.Pass, expr ast.Expr) bool {
	id, ok := ast.Unparen(expr).(*ast.Ident)
	return ok && pass.TypesInfo.ObjectOf(id) == types.Universe.Lookup("nil")
}

func nilSameVariable(pass *analysis.Pass, expr ast.Expr, guard *ast.Ident) bool {
	id, ok := ast.Unparen(expr).(*ast.Ident)
	return ok && pass.TypesInfo.ObjectOf(id) == pass.TypesInfo.ObjectOf(guard)
}

func nilCanBeNil(typ types.Type) bool {
	switch typ.Underlying().(type) {
	case *types.Pointer, *types.Slice, *types.Map, *types.Chan, *types.Signature, *types.Interface:
		return true
	}
	if basic, ok := typ.Underlying().(*types.Basic); ok {
		return basic.Kind() == types.UnsafePointer
	}
	return false
}

// nilSafeAccess handles exactly one guarded selection, optionally called,
// or a direct optional callback. It deliberately avoids rewriting a longer
// primary chain, whose extra guards or boundaries need separate analysis.
func nilSafeAccess(pass *analysis.Pass, expr ast.Expr, guard *ast.Ident) (string, bool) {
	expr = ast.Unparen(expr)
	if selector, ok := expr.(*ast.SelectorExpr); ok && nilSameVariable(pass, selector.X, guard) {
		switch pass.TypesInfo.TypeOf(guard).Underlying().(type) {
		case *types.Pointer, *types.Interface:
			return guard.Name + "?." + selector.Sel.Name, true
		}
	}
	call, ok := expr.(*ast.CallExpr)
	if !ok {
		return "", false
	}
	var fun string
	if nilSameVariable(pass, call.Fun, guard) {
		if _, ok := pass.TypesInfo.TypeOf(guard).Underlying().(*types.Signature); !ok {
			return "", false
		}
		fun = guard.Name + "?"
	} else {
		selector, ok := ast.Unparen(call.Fun).(*ast.SelectorExpr)
		if !ok {
			return "", false
		}
		fun, ok = nilSafeAccess(pass, selector, guard)
		if !ok {
			return "", false
		}
	}
	var args []string
	for _, arg := range call.Args {
		args = append(args, conditionalText(pass, arg))
	}
	arguments := strings.Join(args, ", ")
	if call.Ellipsis.IsValid() {
		arguments += "..."
	}
	return fun + "(" + arguments + ")", true
}

func nilOperand(pass *analysis.Pass, expr ast.Expr) string {
	text := conditionalText(pass, expr)
	// Coalescing must not mix with ordinary binary operators unparenthesized.
	if _, ok := expr.(*ast.BinaryExpr); ok {
		return "(" + text + ")"
	}
	return text
}
