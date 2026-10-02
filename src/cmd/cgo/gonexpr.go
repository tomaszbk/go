//go:build !compiler_bootstrap

package main

import (
	"go/ast"
	"go/token"
)

// walkGonExpr supplies structural traversal for the Gon nodes that are absent
// from the upstream AST used to bootstrap cgo.
func (f *File) walkGonExpr(x any, visit func(*File, any, astContext)) bool {
	switch n := x.(type) {
	case *ast.LambdaExpr:
		// Parameters are names, not type expressions or C references.
		if n.Body != nil {
			f.walk(&n.Body, ctxExpr, visit)
		}
		if n.Block != nil {
			f.walk(n.Block, ctxStmt, visit)
		}
	case *ast.NilGuardExpr:
		f.walk(&n.X, ctxExpr, visit)
	case *ast.SafeNavExpr:
		f.walk(&n.X, ctxExpr, visit)
	default:
		return false
	}
	return true
}

func needsGonTargetType(x ast.Expr) bool {
	switch x := ast.Unparen(x).(type) {
	case *ast.LambdaExpr, *ast.SafeNavExpr:
		return true
	case *ast.BinaryExpr:
		return x.Op == token.COALESCE
	}
	return false
}
