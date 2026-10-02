package astutil

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"testing"
)

func TestGonFeatureRewrites(t *testing.T) {
	const source = `package p
var f func(int) int = (x) => p?.N ?? x
var g func() = () => { callback?() }
`
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "p.go", source, 0)
	if err != nil {
		t.Fatal(err)
	}
	want := map[ast.Node]bool{}
	for n := range ast.Preorder(f) {
		want[n] = true
	}
	Apply(f, func(c *Cursor) bool {
		delete(want, c.Node())
		if id, ok := c.Node().(*ast.Ident); ok && id.Name == "x" {
			c.Replace(&ast.Ident{NamePos: id.NamePos, Name: "renamed"})
		}
		return true
	}, nil)
	if len(want) != 0 {
		t.Fatalf("rewrite omitted %d structural nodes", len(want))
	}
	var formatted bytes.Buffer
	if err := format.Node(&formatted, fs, f); err != nil {
		t.Fatal(err)
	}
	if bytes.Count(formatted.Bytes(), []byte("renamed")) != 2 {
		t.Fatalf("parameter/body rename missing: %s", formatted.Bytes())
	}
	if _, err := parser.ParseFile(token.NewFileSet(), "p.go", formatted.Bytes(), 0); err != nil {
		t.Fatal(err)
	}
	for n := range ast.Preorder(f) {
		switch n.(type) {
		case *ast.LambdaExpr, *ast.NilGuardExpr, *ast.SafeNavExpr:
			if NodeDescription(n) == "" {
				t.Fatalf("missing description for %T", n)
			}
			path, _ := PathEnclosingInterval(f, n.Pos(), n.End())
			if len(path) == 0 {
				t.Fatalf("missing path for %T", n)
			}
		}
	}
}
