// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gonmodernize

import (
	"go/ast"
	"go/format"
	"go/token"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
)

// ConditionalAnalyzer suggests conditional expressions for simple value choices.
var ConditionalAnalyzer = &analysis.Analyzer{
	Name: "gonconditional",
	Doc: `replace simple if/else value choices with Gon conditional expressions

Replaces if/else branches that each return one value or assign one local
variable. Branches retain their original target types and lazy evaluation.
Statements with initializers, comments, multi-value results or complex assignment
targets are left unchanged. More specific nil
operators take precedence over conditional expressions.`,
	Run: runConditional,
}

func runConditional(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		if ast.IsGenerated(file) {
			continue
		}
		conditionalWalk(pass, file, nil, func(stmt *ast.IfStmt, sig *types.Signature) bool {
			if conditionalHasComments(file, stmt) {
				return false
			}
			choice, ok := conditionalChoice(pass, stmt, sig)
			if !ok {
				return false
			}
			// Prefer the more specific nil operators to a conditional expression.
			if _, ok := nilOpportunity(pass, stmt, sig); ok {
				return true
			}
			replacement := choice.prefix + "if " + conditionalText(pass, stmt.Cond) + " { " +
				conditionalText(pass, choice.then) + " } else { " + conditionalText(pass, choice.otherwise) + " }"
			conditionalReport(pass, stmt, "if/else value choice can use a conditional expression", "Use conditional expression", replacement)
			return true
		})
	}
	return nil, nil
}

type conditionalValueChoice struct {
	prefix          string
	target          types.Type
	then, otherwise ast.Expr
}

// conditionalChoice accepts only contexts that preserve the target type of
// each branch. A local identifier assignment has no LHS evaluation to move
// before the condition; indexed and indirect assignments do.
func conditionalChoice(pass *analysis.Pass, stmt *ast.IfStmt, sig *types.Signature) (conditionalValueChoice, bool) {
	var result conditionalValueChoice
	other, ok := stmt.Else.(*ast.BlockStmt)
	if stmt.Init != nil || !ok || len(stmt.Body.List) != 1 || len(other.List) != 1 {
		return result, false
	}
	switch then := stmt.Body.List[0].(type) {
	case *ast.ReturnStmt:
		otherwise, ok := other.List[0].(*ast.ReturnStmt)
		if !ok || len(then.Results) != 1 || len(otherwise.Results) != 1 || sig == nil || sig.Results().Len() != 1 {
			return result, false
		}
		result = conditionalValueChoice{"return ", sig.Results().At(0).Type(), then.Results[0], otherwise.Results[0]}
	case *ast.AssignStmt:
		otherwise, ok := other.List[0].(*ast.AssignStmt)
		if !ok || then.Tok != token.ASSIGN || otherwise.Tok != token.ASSIGN ||
			len(then.Lhs) != 1 || len(then.Rhs) != 1 || len(otherwise.Lhs) != 1 || len(otherwise.Rhs) != 1 {
			return result, false
		}
		variable := conditionalLocal(pass, then.Lhs[0])
		if variable == nil || variable != conditionalLocal(pass, otherwise.Lhs[0]) {
			return result, false
		}
		result = conditionalValueChoice{conditionalText(pass, then.Lhs[0]) + " = ", variable.Type(), then.Rhs[0], otherwise.Rhs[0]}
	default:
		return result, false
	}
	for _, expr := range []ast.Expr{result.then, result.otherwise} {
		if typ := pass.TypesInfo.TypeOf(expr); typ == nil {
			return result, false
		} else if _, tuple := typ.(*types.Tuple); tuple {
			return result, false
		}
		// Direct nested conditionals are intentionally not part of the language.
		if _, nested := ast.Unparen(expr).(*ast.CondExpr); nested {
			return result, false
		}
	}
	if !conditionalInterfaceSafe(pass, result) {
		return result, false
	}
	return result, true
}

// Conditional/coalescing expressions combine the no-target types of untyped
// branches before boxing them in an interface. Two independent assignments
// instead default each constant separately. Recheck without a target because
// the original TypesInfo has already recorded those assignment default types.
func conditionalInterfaceSafe(pass *analysis.Pass, choice conditionalValueChoice) bool {
	if _, parameter := types.Unalias(choice.target).(*types.TypeParam); parameter {
		return true
	}
	if _, iface := choice.target.Underlying().(*types.Interface); !iface {
		return true
	}
	typesWithoutTarget := make([]types.Type, 2)
	for i, expr := range []ast.Expr{choice.then, choice.otherwise} {
		info := &types.Info{Types: make(map[ast.Expr]types.TypeAndValue)}
		if err := types.CheckExpr(pass.Fset, pass.Pkg, expr.Pos(), expr, info); err != nil {
			return false
		}
		typesWithoutTarget[i] = info.TypeOf(expr)
	}
	for i, typ := range typesWithoutTarget {
		basic, ok := typ.(*types.Basic)
		if !ok || basic.Info()&types.IsUntyped == 0 || basic.Kind() == types.UntypedNil {
			continue
		}
		other := typesWithoutTarget[1-i]
		if _, iface := other.Underlying().(*types.Interface); iface {
			continue
		}
		if basic, ok := other.(*types.Basic); ok && basic.Kind() == types.UntypedNil {
			continue
		}
		if !types.Identical(types.Default(typ), types.Default(other)) {
			return false
		}
	}
	return true
}

func conditionalLocal(pass *analysis.Pass, expr ast.Expr) *types.Var {
	id, ok := ast.Unparen(expr).(*ast.Ident)
	if !ok {
		return nil
	}
	variable, ok := pass.TypesInfo.ObjectOf(id).(*types.Var)
	if !ok || variable.IsField() || variable.Parent() == nil || variable.Parent() == pass.Pkg.Scope() {
		return nil
	}
	return variable
}

// conditionalWalk tracks function boundaries explicitly, including Gon lambdas.
// Once a replacement is offered, do not offer edits inside its replaced range.
func conditionalWalk(pass *analysis.Pass, node ast.Node, sig *types.Signature, visit func(*ast.IfStmt, *types.Signature) bool) {
	switch node := node.(type) {
	case *ast.FuncDecl:
		sig, _ = pass.TypesInfo.TypeOf(node.Name).(*types.Signature)
	case *ast.FuncLit:
		sig, _ = pass.TypesInfo.TypeOf(node).(*types.Signature)
	case *ast.LambdaExpr:
		sig, _ = pass.TypesInfo.TypeOf(node).(*types.Signature)
	case *ast.IfStmt:
		if visit(node, sig) {
			return
		}
	}
	for child := range ast.Children(node) {
		conditionalWalk(pass, child, sig, visit)
	}
}

func conditionalHasComments(file *ast.File, node ast.Node) bool {
	for _, comment := range file.Comments {
		if comment.Pos() < node.End() && comment.End() > node.Pos() {
			return true
		}
	}
	return false
}

func conditionalText(pass *analysis.Pass, node ast.Node) string {
	var out strings.Builder
	format.Node(&out, pass.Fset, node)
	return out.String()
}

func conditionalReport(pass *analysis.Pass, stmt *ast.IfStmt, message, fix, replacement string) {
	pass.Report(analysis.Diagnostic{
		Pos: stmt.Pos(), End: stmt.Cond.End(), Message: message,
		SuggestedFixes: []analysis.SuggestedFix{{
			Message:   fix,
			TextEdits: []analysis.TextEdit{{Pos: stmt.Pos(), End: stmt.End(), NewText: []byte(replacement)}},
		}},
	})
}
