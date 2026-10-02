//go:build !compiler_bootstrap

package main

import "go/ast"

// Bootstrap cgo processes only standard Go source and is built against the
// upstream go/ast. condexpr_bootstrap.go supplies its no-op adapters.
func (f *File) walkCondExpr(x any, visit func(*File, any, astContext)) bool {
	n, ok := x.(*ast.CondExpr)
	if !ok {
		return false
	}
	f.walk(&n.Cond, ctxExpr, visit)
	f.walk(&n.Then, ctxExpr, visit)
	f.walk(&n.Else, ctxExpr, visit)
	return true
}

// isCondExpr reports whether x requires the C parameter's target type when
// cgo moves it into a temporary, including through parentheses.
func isCondExpr(x ast.Expr) bool {
	for {
		p, ok := x.(*ast.ParenExpr)
		if !ok {
			break
		}
		x = p.X
	}
	_, ok := x.(*ast.CondExpr)
	return ok
}
