//go:build compiler_bootstrap

package main

import "go/ast"

// See condexpr.go.
func (f *File) walkCondExpr(x any, visit func(*File, any, astContext)) bool {
	return false
}

func isCondExpr(x ast.Expr) bool {
	return false
}
