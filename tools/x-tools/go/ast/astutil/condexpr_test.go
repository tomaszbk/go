package astutil

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"slices"
	"testing"
)

func findCondExpr(n ast.Node) *ast.CondExpr {
	var cond *ast.CondExpr
	ast.Inspect(n, func(n ast.Node) bool {
		if c, ok := n.(*ast.CondExpr); ok && cond == nil {
			cond = c
		}
		return cond == nil
	})
	return cond
}

func TestCondExpr(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "p.go", "package p\n\nvar x = f(if c { a } else { b })\n", 0)
	if err != nil {
		t.Fatal(err)
	}
	cond := findCondExpr(f)

	// The tokens of a conditional expression belong to it.
	for _, tok := range []struct {
		pos token.Pos
		len int
	}{
		{cond.If, len("if")},
		{cond.Lbrace, len("{")},
		{cond.Rbrace, len("}")},
		{cond.ElsePos, len("else")},
		{cond.ElseLbrace, len("{")},
		{cond.ElseRbrace, len("}")},
	} {
		path, _ := PathEnclosingInterval(f, tok.pos, tok.pos+token.Pos(tok.len))
		if path[0] != cond {
			t.Errorf("token at %s: innermost node is %T, want the conditional expression", fset.Position(tok.pos), path[0])
		}
	}
	for _, operand := range []ast.Expr{cond.Cond, cond.Then, cond.Else} {
		path, exact := PathEnclosingInterval(f, operand.Pos(), operand.End())
		if path[0] != operand || path[1] != cond || !exact {
			t.Errorf("operand at %s: path %T, %T (exact=%v), want the operand in the conditional expression", fset.Position(operand.Pos()), path[0], path[1], exact)
		}
	}
	if path, exact := PathEnclosingInterval(f, cond.Pos(), cond.End()); path[0] != cond || !exact {
		t.Errorf("whole expression: innermost node is %T (exact=%v)", path[0], exact)
	}
	if got := NodeDescription(cond); got != "conditional expression" {
		t.Errorf("NodeDescription = %q", got)
	}

	// Apply visits the condition and the branches, which can be replaced.
	var names []string
	Apply(f, func(c *Cursor) bool {
		if c.Parent() == cond {
			names = append(names, c.Name())
		}
		if id, ok := c.Node().(*ast.Ident); ok && id.Name == "b" {
			c.Replace(ast.NewIdent("z"))
		}
		return true
	}, nil)
	if want := []string{"Cond", "Then", "Else"}; !slices.Equal(names, want) {
		t.Errorf("Apply visited %q of the conditional expression, want %q", names, want)
	}
	var buf bytes.Buffer
	if err := format.Node(&buf, fset, f); err != nil {
		t.Fatal(err)
	}
	if want := "package p\n\nvar x = f(if c { a } else { z })\n"; buf.String() != want {
		t.Errorf("after Apply:\n%s\nwant:\n%s", buf.String(), want)
	}
}

// A conditional expression recovered from a syntax error may lack the
// positions of its else branch.
func TestCondExprMissingElse(t *testing.T) {
	const src = "package p\n\nvar y = g(if c { a }, d)\n"
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "q.go", src, parser.AllErrors)
	if err == nil {
		t.Fatal("no syntax error")
	}
	cond := findCondExpr(f)
	if cond == nil || cond.ElsePos.IsValid() {
		t.Fatalf("parser did not recover a conditional expression without else: %#v", cond)
	}
	file := fset.File(f.Pos())
	for off := 0; off < len(src); off++ {
		pos := file.Pos(off)
		if path, _ := PathEnclosingInterval(f, pos, pos+1); len(path) == 0 {
			t.Errorf("offset %d: empty path", off)
		}
	}
	if path, _ := PathEnclosingInterval(f, cond.If, cond.If+2); path[0] != cond {
		t.Errorf("if token: innermost node is %T, want the conditional expression", path[0])
	}
}
