//go:build compiler_bootstrap

package main

import "go/ast"

// See errorexpr.go.

func (f *File) walkErrorExpr(x any, visit func(*File, any, astContext)) bool {
	return false
}

func findErrorExpr(x ast.Expr) ast.Node {
	return nil
}
