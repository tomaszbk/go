package printer

import (
	"go/ast"
	"reflect"
)

var (
	exprType      = reflect.TypeFor[ast.Expr]()
	exprSliceType = reflect.TypeFor[[]ast.Expr]()
)

// stripCondParens removes all parentheses around conditional expressions in
// f, as in syntax trees built by programs rather than the parser. Where they
// are needed, at the start of a statement, the printer must restore them.
func stripCondParens(f *ast.File) {
	unparen := func(v reflect.Value) {
		if x, ok := v.Interface().(*ast.ParenExpr); ok {
			if c, ok := ast.Unparen(x).(*ast.CondExpr); ok {
				v.Set(reflect.ValueOf(c))
			}
		}
	}
	ast.Inspect(f, func(n ast.Node) bool {
		v := reflect.ValueOf(n)
		if v.Kind() != reflect.Pointer || v.Elem().Kind() != reflect.Struct {
			return true
		}
		v = v.Elem()
		for i := range v.NumField() {
			switch f := v.Field(i); f.Type() {
			case exprType:
				if !f.IsNil() {
					unparen(f)
				}
			case exprSliceType:
				for j := range f.Len() {
					unparen(f.Index(j))
				}
			}
		}
		return true
	})
}
