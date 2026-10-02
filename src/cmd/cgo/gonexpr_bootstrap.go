//go:build compiler_bootstrap

package main

import "go/ast"

func (f *File) walkGonExpr(x any, visit func(*File, any, astContext)) bool {
	return false
}

func needsGonTargetType(x ast.Expr) bool {
	return false
}
