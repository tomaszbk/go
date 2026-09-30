package astutil

import "go/ast"

// Unparen returns e with any enclosing parentheses stripped.
// Deprecated: use [ast.Unparen].
//
//go:fix inline
func Unparen(e ast.Expr) ast.Expr { return ast.Unparen(e) }
