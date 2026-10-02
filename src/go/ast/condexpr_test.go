package ast_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"slices"
	"strings"
	"testing"
)

func TestCondExprWalk(t *testing.T) {
	const src = `package p; var x = if a { b } else { c(d) }`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, parser.SkipObjectResolution)
	if err != nil {
		t.Fatal(err)
	}
	var cond *ast.CondExpr
	var events []string
	for n := range ast.Preorder(f) {
		if x, ok := n.(*ast.CondExpr); ok {
			cond = x
		}
		if cond != nil {
			event := strings.TrimPrefix(reflect.TypeOf(n).String(), "*ast.")
			if id, ok := n.(*ast.Ident); ok {
				event += " " + id.Name
			}
			events = append(events, event)
		}
	}
	want := []string{"CondExpr", "Ident a", "Ident b", "CallExpr", "Ident c", "Ident d"}
	if !slices.Equal(events, want) {
		t.Errorf("Preorder events:\ngot:  %s\nwant: %s", events, want)
	}
	if cond == nil {
		t.Fatal("no conditional expression")
	}
	if got, want := src[fset.Position(cond.Pos()).Offset:fset.Position(cond.End()).Offset], "if a { b } else { c(d) }"; got != want {
		t.Errorf("source range of conditional expression = %q, want %q", got, want)
	}
}

func TestCondExprPosEnd(t *testing.T) {
	x := &ast.CondExpr{If: 10, Lbrace: 15, Rbrace: 19, ElsePos: 21, ElseLbrace: 26, ElseRbrace: 30}
	if x.Pos() != 10 || x.End() != 31 {
		t.Errorf("Pos(), End() = %d, %d, want 10, 31", x.Pos(), x.End())
	}
	var _ ast.Expr = x // CondExpr is an expression
}
